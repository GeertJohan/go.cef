package main

import "github.com/GeertJohan/go.cef"

func main() {
	// create fake args
	args := make([]string, 0)

	// // Provide CEF with command-line arguments.
	// CefMainArgs main_args(argc, argv);
	//++ TODO?

	// // SimpleApp implements application-level callbacks. It will create the first
	// // browser instance in OnContextInitialized() after CEF has initialized.
	// CefRefPtr<SimpleApp> app(new SimpleApp);
	app := cef.NewApp(cef.AppHandler)

	// // CEF applications have multiple sub-processes (render, plugin, GPU, etc)
	// // that share the same executable. This function checks the command-line and,
	// // if this is a sub-process, executes the appropriate logic.
	// int exit_code = CefExecuteProcess(main_args, app.get(), NULL);
	// if (exit_code >= 0) {
	// 	// The sub-process has completed so return here.
	// 	return exit_code;
	// }
	err := cef.ExecuteProcess(args, app, nil)
	if err != nil {
		eerr := err.(cef.CodeError)
	}

	// // Specify CEF global settings here.
	// CefSettings settings;
	//++ TODO Global cef.Settings ??

	// // Initialize CEF for the browser process.
	// CefInitialize(main_args, settings, app.get(), NULL);
	cef.Initialize(args, settings, app, nil)

	// // Run the CEF message loop. This will block until CefQuitMessageLoop() is
	// // called.
	// CefRunMessageLoop();
	cef.RunMessageLoop()

	// // Shut down CEF.
	// CefShutdown();
	cef.Shutdown()
}
