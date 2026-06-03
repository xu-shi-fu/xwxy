package main4httppingclient
import (
    p0ef6f2938 "github.com/starter-go/application"
    pc8016548a "github.com/xu-shi-fu/xwxy/tools/httpping/app/client"
     "github.com/starter-go/application"
)

// type pc8016548a.Runner in package:github.com/xu-shi-fu/xwxy/tools/httpping/app/client
//
// id:com-c8016548a1981c33-client-Runner
// class:
// alias:
// scope:singleton
//
type pc8016548a1_client_Runner struct {
}

func (inst* pc8016548a1_client_Runner) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-c8016548a1981c33-client-Runner"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pc8016548a1_client_Runner) new() any {
    return &pc8016548a.Runner{}
}

func (inst* pc8016548a1_client_Runner) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pc8016548a.Runner)
	nop(ie, com)

	
    com.URL = inst.getURL(ie)
    com.IntervalMs = inst.getIntervalMs(ie)
    com.AC = inst.getAC(ie)


    return nil
}


func (inst*pc8016548a1_client_Runner) getURL(ie application.InjectionExt)string{
    return ie.GetString("${httpping.client.location}")
}


func (inst*pc8016548a1_client_Runner) getIntervalMs(ie application.InjectionExt)int64{
    return ie.GetInt64("${httpping.client.interval}")
}


func (inst*pc8016548a1_client_Runner) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


