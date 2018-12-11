package handlepanic

import "sync"

type panicHandler struct {
	actualHandlePanicFunction func()
}

var instance *panicHandler
var confOnce sync.Once

//CreateInstance function is to create an instance
func CreateInstance(handler func()) {
	confOnce.Do(func() {
		instance = &panicHandler{actualHandlePanicFunction: handler}
	})
}

//Handle function to handle panic
//fn is a generic function, ideally it does not have any return type as its being called as a go routine and hence its return type is ignored anyways
func Handle(fn func(interface{}) interface{}, args interface{}) {
	instance.handle(fn, args)
}

func (h panicHandler) handle(fn func(interface{}) interface{}, args interface{}) {
	defer h.actualHandlePanicFunction()
	fn(args)
}
