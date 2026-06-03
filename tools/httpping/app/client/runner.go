package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"

	"github.com/starter-go/application"
	"github.com/starter-go/vlog"
)

type Runner struct {

	//starter:component

	URL        string //starter:inject("${httpping.client.location}")
	IntervalMs int64  //starter:inject("${httpping.client.interval}")

	AC application.Context //starter:inject("context")

	stopping bool
	tracks   map[TrackURL]*track
}

// Life implements application.Lifecycle.
func (inst *Runner) Life() *application.Life {
	l := &application.Life{
		OnCreate: inst.onCreate,
		OnLoop:   inst.run,
	}
	return l
}

func (inst *Runner) onCreate() error {

	err := inst.innerLoadTracks()
	if err != nil {
		return err
	}

	return inst.innerInitBeep()
}

func (inst *Runner) run() error {

	vlog.Info("on:client.Runner.run()")

	ms := inst.IntervalMs
	interval := time.Millisecond * time.Duration(ms)

	for {

		if inst.stopping {
			break
		}

		err := inst.runOnce()
		if err != nil {
			vlog.Warn("%s", err.Error())
		}

		time.Sleep(interval)
	}

	return nil
}

func (inst *Runner) innerLoadTracks() error {
	refs := []TrackURL{
		TrackPing,
		TrackPong,
		TrackError,
	}
	table := make(map[TrackURL]*track)
	for _, url := range refs {
		tr, err := inst.innerLoadTrack(url)
		if err != nil {
			return err
		}
		table[url] = tr
	}

	inst.tracks = table
	return nil
}

func (inst *Runner) innerInitBeep() error {

	// 初始化扬声器

	tr := inst.innerGetTrackRequired(TrackPing)
	format := tr.format
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	return nil
}

func (inst *Runner) innerClose(c io.Closer) {
	if c == nil {
		return
	}
	c.Close()
}

func (inst *Runner) innerLoadTrack(url TrackURL) (*track, error) {

	ref := string(url)
	resSet := inst.AC.GetResources()

	src, err := resSet.Open(ref)
	if err != nil {
		return nil, err
	}
	defer inst.innerClose(src)

	// 解码
	streamer, format, err := mp3.Decode(src)
	if err != nil {
		return nil, err
	}
	defer inst.innerClose(streamer)

	tr := new(track)
	tr.url = url
	tr.format = format
	tr.stream = streamer

	streamer = nil

	return tr, nil
}

func (inst *Runner) innerGetTrackRequired(url TrackURL) *track {
	all := inst.tracks
	tr := all[url]
	if tr == nil {
		panic("no Sound-Track with URL = " + url)
	}
	return tr
}

func (inst *Runner) innerPlaySound(t *task, url TrackURL) {

	sel := t.trackCurrent
	switch url {
	case TrackPing:
		sel = t.trackPing
	case TrackPong:
		sel = t.trackPong
	case TrackError:
		sel = t.trackError
	}
	t.trackCurrent = sel
	inst.innerDoPlaySound(t)
}

func (inst *Runner) runOnce() error {

	t := new(task)
	steps := make([]func(*task) error, 0)

	steps = append(steps, inst.innerDoInit)
	steps = append(steps, inst.innerDoSetupTracks)
	steps = append(steps, inst.innerDoPrepareRequest)
	steps = append(steps, inst.innerDoSend)
	steps = append(steps, inst.innerDoHandleResponse)
	steps = append(steps, inst.innerDoCheckResult)
	steps = append(steps, inst.innerDoLogResult)

	// steps = append(steps, inst.innerDoPlaySound)
	// steps = append(steps, inst.innerDoPlaySound)

	for _, st := range steps {
		err := st(t)
		if err != nil {
			inst.innerPlaySound(t, TrackError)
			return err
		}
	}

	return nil
}

func (inst *Runner) innerDoInit(t *task) error {

	t.context = context.Background()
	t.t0 = time.Now()
	t.url = inst.URL
	t.agent = http.DefaultClient
	t.trackCurrent = t.trackPing

	return nil
}

func (inst *Runner) innerDoSetupTracks(t *task) error {

	t.trackPing = inst.innerGetTrackRequired(TrackPing)
	t.trackPong = inst.innerGetTrackRequired(TrackPong)
	t.trackError = inst.innerGetTrackRequired(TrackError)

	t.trackCurrent = t.trackPing

	return nil
}

func (inst *Runner) innerDoPrepareRequest(t *task) error {
	method := http.MethodGet
	url := inst.URL
	req, err := http.NewRequest(method, url, nil)
	if err == nil {
		t.req = req
	}

	inst.innerPlaySound(t, TrackPing)

	return err
}

func (inst *Runner) innerDoSend(t *task) error {
	req := t.req
	client := t.agent
	resp, err := client.Do(req)
	if err == nil {
		t.resp = resp
	}
	return err
}

func (inst *Runner) innerDoHandleResponse(t *task) error {

	resp := t.resp
	if resp == nil {
		return fmt.Errorf("response is nil")
	}

	code := resp.StatusCode
	if code != http.StatusOK {
		return fmt.Errorf("HTTP %s", resp.Status)
	}

	body := resp.Body
	if body == nil {
		return fmt.Errorf("response body is nil")
	}
	defer inst.innerClose(body)

	dst := new(nopWriter)
	dst.init()
	inst.innerClose(dst)
	io.Copy(dst, body)
	return nil
}

func (inst *Runner) innerDoCheckResult(t *task) error {

	resp := t.resp
	if resp == nil {
		return fmt.Errorf("response is nil")
	}

	code := resp.StatusCode
	if code != http.StatusOK {
		return fmt.Errorf("HTTP %s", resp.Status)
	}

	t.t1 = time.Now()

	return nil
}

func (inst *Runner) innerDoLogResult(t *task) error {

	req := t.req
	resp := t.resp
	method := req.Method
	url := req.URL.String()
	status := resp.Status

	vlog.Info("%s %s >>> %s", method, url, status)

	inst.innerPlaySound(t, TrackPong)

	return nil

}

func (inst *Runner) innerDoPlaySound(t *task) error {

	// 播放

	current := t.trackCurrent
	if current == nil {
		return nil
	}

	streamer := current.stream
	streamer.Seek(0)

	// vlog.Warn("todo: play sound (%s)", current.url)

	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		t.trackCurrent = nil
	})))

	const step = time.Second / 10
	for ttl := time.Second * 3; ttl > 0; {
		if t.trackCurrent == nil {
			break
		}
		time.Sleep(step)
		ttl -= step
	}

	return nil
}

func (inst *Runner) _impl() application.Lifecycle {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type TrackURL string

const (
	TrackPing  TrackURL = "res:///sounds/ping.mp3"
	TrackPong  TrackURL = "res:///sounds/pong.mp3"
	TrackError TrackURL = "res:///sounds/error.mp3"
)

////////////////////////////////////////////////////////////////////////////////

type task struct {
	context context.Context
	t0      time.Time
	t1      time.Time
	url     string
	req     *http.Request
	resp    *http.Response
	agent   *http.Client
	err     error

	// tracks

	trackPing    *track
	trackPong    *track
	trackError   *track
	trackCurrent *track
}

////////////////////////////////////////////////////////////////////////////////

type track struct {
	url    TrackURL
	format beep.Format
	stream beep.StreamSeekCloser
}

////////////////////////////////////////////////////////////////////////////////

type nopWriter struct {
	count int64
	gate  int64
}

// Close implements io.Closer.
func (inst *nopWriter) Close() error {

	inst.gate = 0
	inst.tryLog()

	return nil
}

// Write implements io.Writer.
func (inst *nopWriter) Write(p []byte) (n int, err error) {

	size := len(p)
	inst.count += int64(size)

	inst.tryLog()

	return size, nil
}

func (inst *nopWriter) init() {
	inst.count = 0
	inst.gate = 16
}

func (inst *nopWriter) tryLog() {

	if inst.count < inst.gate {
		return
	}

	inst.updateGate()

	vlog.Info("rx: %d (bytes)", inst.count)

}

func (inst *nopWriter) updateGate() {

	const limit = 1024 * 1024 * 32
	g := inst.gate

	if g < 1 {
		g = 1
	} else if g < limit {
		g = g * 2
	} else {
		g = g + limit
	}

	inst.gate = g
}

////////////////////////////////////////////////////////////////////////////////
