package cef

// AppHandler interface specifies the hooks that can be called
// TODO: not sure if using interface is correct.
//		It forces one to implement ALL functions even though that might not be required.
//		Maybe just have the On* function hooks as fields on the App struct and drop the interface
type AppHandler interface {
	OnDoSomething
	onDoBar
}

// OnDoFoo is called when foo must be done
type OnDoFoo func() bool

// OnDoBar is called when bar must be done
type OnDoBar func(string) error

// App instance
type App struct {
	AppHandler
	// app instance
}
