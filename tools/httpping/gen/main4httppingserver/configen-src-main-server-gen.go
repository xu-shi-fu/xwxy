package main4httppingserver
import (
    pd1a916a20 "github.com/starter-go/libgin"
    p6a7ddd006 "github.com/xu-shi-fu/xwxy/tools/httpping/app/server/web/controllers"
     "github.com/starter-go/application"
)

// type p6a7ddd006.ExampleController in package:github.com/xu-shi-fu/xwxy/tools/httpping/app/server/web/controllers
//
// id:com-6a7ddd006204be00-controllers-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p6a7ddd0062_controllers_ExampleController struct {
}

func (inst* p6a7ddd0062_controllers_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6a7ddd006204be00-controllers-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6a7ddd0062_controllers_ExampleController) new() any {
    return &p6a7ddd006.ExampleController{}
}

func (inst* p6a7ddd0062_controllers_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6a7ddd006.ExampleController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)


    return nil
}


func (inst*p6a7ddd0062_controllers_ExampleController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}



// type p6a7ddd006.PingController in package:github.com/xu-shi-fu/xwxy/tools/httpping/app/server/web/controllers
//
// id:com-6a7ddd006204be00-controllers-PingController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p6a7ddd0062_controllers_PingController struct {
}

func (inst* p6a7ddd0062_controllers_PingController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6a7ddd006204be00-controllers-PingController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6a7ddd0062_controllers_PingController) new() any {
    return &p6a7ddd006.PingController{}
}

func (inst* p6a7ddd0062_controllers_PingController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6a7ddd006.PingController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)


    return nil
}


func (inst*p6a7ddd0062_controllers_PingController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


