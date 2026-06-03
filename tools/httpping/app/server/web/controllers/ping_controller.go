package controllers

import (
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/libgin"
	"github.com/starter-go/vlog"
	"github.com/xu-shi-fu/xwxy/tools/httpping/app/common/data/dxo"
	"github.com/xu-shi-fu/xwxy/tools/httpping/app/server/web/dto"
	"github.com/xu-shi-fu/xwxy/tools/httpping/app/server/web/vo"
)

////////////////////////////////////////////////////////////////////////////////

type PingController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Sender libgin.Responder //starter:inject("#")
}

func (inst *PingController) _impl() libgin.Controller {
	return inst
}

func (inst *PingController) Registration() *libgin.ControllerRegistration {
	cr1 := new(libgin.ControllerRegistration)
	cr1.Route = inst.route
	return cr1
}

func (inst *PingController) route(rp libgin.RouterProxy) error {

	rp = rp.For("ping")

	rp.GET("", inst.handleGetList)
	rp.GET("long-fetch/:filename", inst.handleGetLongFetch)

	// rp.GET(":id", inst.handleGetOne)
	// rp.PUT(":id", inst.handlePutItem)

	return nil
}

func (inst *PingController) handleGetOne(gc *gin.Context) {

	req := new(myPingRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetOne)
}

func (inst *PingController) handleGetList(gc *gin.Context) {

	req := new(myPingRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetList)
}

func (inst *PingController) handleGetLongFetch(gc *gin.Context) {

	req := new(myPingRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetLongFetch)
}

func (inst *PingController) handlePutItem(gc *gin.Context) {

	req := new(myPingRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = true
	req.wantRequestBody = true

	req.execute(req.doPutItem)
}

////////////////////////////////////////////////////////////////////////////////

type myPingRequest struct {
	wantRequestID    bool
	wantRequestBody  bool
	skipResponseBody bool

	context    *gin.Context
	controller *PingController

	id    dxo.PingID
	body1 vo.Pings
	body2 vo.Pings
}

func (inst *myPingRequest) open(ctx *gin.Context) error {

	if inst.wantRequestID {
		str := ctx.Param("id")
		num, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		inst.id = dxo.PingID(num)
	}

	if inst.wantRequestBody {
		obj := &inst.body1
		err := ctx.BindJSON(obj)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myPingRequest) execute(task func() error) {

	// ex := new(libgin.Executor)
	// ex.Context = inst.context
	// ex.Responder = inst.controller.Sender
	// ex.Body1 = &inst.body1
	// ex.Body2 = &inst.body2
	// ex.OnOpen = inst.open
	// ex.OnTask = task
	// ex.Execute()

	ctx := inst.context
	err := inst.open(ctx)
	if err == nil {
		err = task()
	}
	inst.send(err)
}

func (inst *myPingRequest) send(err error) {

	if inst.skipResponseBody {
		return
	}

	ctx := inst.context
	status := inst.body2.Status
	data := &inst.body2

	resp := &libgin.Response{
		Context: ctx,
		Status:  status,
		Error:   err,
		Data:    data,
	}
	inst.controller.Sender.Send(resp)
}

func (inst *myPingRequest) doGetList() error {

	it := &dto.Ping{}
	inst.body2.Items = []*dto.Ping{it, it, it}
	return nil
}

func (inst *myPingRequest) doGetLongFetch() error {

	const (
		bitSize = 64
		base    = 10
	)

	// query

	c := inst.context
	strLength := c.Query("length")
	strBps := c.Query("bps")
	nLen, err1 := strconv.ParseInt(strLength, base, bitSize)
	nBps, err2 := strconv.ParseInt(strBps, base, bitSize)
	if err1 != nil {
		vlog.Warn("want query['length']")
		return err1
	}
	if err2 != nil {
		vlog.Warn("want query['bps']")
		return err2
	}

	// send

	inst.skipResponseBody = true

	code := http.StatusOK
	ctype := "application/x-bin-stream"
	src := &myMockDataStream{
		length: nLen,
		bps:    nBps,
	}
	src.init()

	c.DataFromReader(code, nLen, ctype, src, nil)

	return nil
}

func (inst *myPingRequest) doGetOne() error {

	it := &dto.Ping{}

	inst.body2.Items = []*dto.Ping{it}
	return nil
}

func (inst *myPingRequest) doPutItem() error {

	it1 := inst.body1.Items[0]
	it2 := &dto.Ping{}
	id := inst.id

	it2.ID = id

	inst.body2.Items = []*dto.Ping{it1, it2}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type myMockDataStream struct {
	bps      int64 // limit
	length   int64 // total
	count    int64 // done
	t0       lang.Time
	template []byte
}

func (inst *myMockDataStream) init() {

	now := lang.Now()

	inst.t0 = now
	inst.count = 0
	inst.template = inst.makeTemplate()

}

func (inst *myMockDataStream) makeTemplate() []byte {

	b := new(strings.Builder)

	for c := 'a'; c <= 'z'; c++ {
		b.WriteRune(c)
	}
	for c := '0'; c <= '9'; c++ {
		b.WriteRune(c)
	}
	for c := 'A'; c <= 'Z'; c++ {
		b.WriteRune(c)
	}

	str := b.String()
	return []byte(str)
}

func (inst *myMockDataStream) checkSpeed() bool {

	const (
		min  = 0.01
		step = time.Millisecond * 10
	)

	now := lang.Now()
	t0 := inst.t0
	cb := inst.count
	sec := float32(now-t0) / 1000

	if sec < min {
		sec = min
	}

	bpsWant := float32(inst.bps)
	bpsHave := float32(cb*8) / (sec)
	ok := true

	if bpsHave > bpsWant {
		// limit the speed
		time.Sleep(step)
		ok = false
	}

	return ok
}

// Read implements io.Reader.
func (inst *myMockDataStream) Read(p []byte) (n int, err error) {

	if inst.length <= inst.count {
		return 0, io.EOF
	}

	bufferlen := len(p)
	rem := inst.length - inst.count // 剩余 bytes
	cb := 0

	if rem < int64(bufferlen) {
		cb = int(rem)
	} else {
		cb = bufferlen
	}

	inst.count += int64(cb)

	for {
		ok := inst.checkSpeed()
		if ok {
			break
		}
	}

	templ := inst.template
	templen := len(templ)
	for i := 0; i < cb; i++ {
		ch := templ[i%templen]
		if i%80 == 0 {
			ch = '\n'
		}
		p[i] = ch
	}

	return cb, nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
