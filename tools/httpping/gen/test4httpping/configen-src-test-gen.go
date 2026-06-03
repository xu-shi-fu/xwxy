package test4httpping
import (
    pef2c32263 "github.com/xu-shi-fu/xwxy/tools/httpping/src/test/golang/unittestcases"
     "github.com/starter-go/application"
)

// type pef2c32263.ExampleCase in package:github.com/xu-shi-fu/xwxy/tools/httpping/src/test/golang/unittestcases
//
// id:com-ef2c32263ab67783-unittestcases-ExampleCase
// class:
// alias:
// scope:singleton
//
type pef2c32263a_unittestcases_ExampleCase struct {
}

func (inst* pef2c32263a_unittestcases_ExampleCase) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ef2c32263ab67783-unittestcases-ExampleCase"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pef2c32263a_unittestcases_ExampleCase) new() any {
    return &pef2c32263.ExampleCase{}
}

func (inst* pef2c32263a_unittestcases_ExampleCase) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pef2c32263.ExampleCase)
	nop(ie, com)

	


    return nil
}


