package cef

// required:
// sudo aptitude install libgtk2.0-dev

/*
#cgo CFLAGS: -I/home/geertjohan/Build/cef_binary_3.1768.1579_linux64
#cgo CFLAGS: -I/usr/include/gtk-2.0 -I/usr/include/glib-2.0
#cgo CFLAGS: -I/usr/lib/x86_64-linux-gnu/glib-2.0/include
#cgo CFLAGS: -I/usr/include/cairo
#cgo CFLAGS: -I/usr/include/pango-1.0
#cgo CFLAGS: -I/usr/lib/x86_64-linux-gnu/gtk-2.0/include
#cgo CFLAGS: -I/usr/include/gdk-pixbuf-2.0
#cgo CFLAGS: -I/usr/include/atk-1.0
#cgo LDFLAGS: -L/home/geertjohan/Build/cef_binary_3.1768.1571_linux64
#include "/home/geertjohan/Build/cef_binary_3.1768.1579_linux64/include/capi/cef_app_capi.h"

#include <stdlib.h>
*/
import "C"


///
// This function should be called from the application entry point function to
// execute a secondary process. It can be used to run secondary processes from
// the browser client executable (default behavior) or from a separate
// executable specified by the CefSettings.browser_subprocess_path value. If
// called for the browser process (identified by no "type" command-line value)
// it will return immediately with a value of -1. If called for a recognized
// secondary process it will block until the process should exit and then return
// the process exit code. The |application| parameter may be NULL. The
// |windows_sandbox_info| parameter is only used on Windows and may be NULL (see
// cef_sandbox_win.h for details).
///
CEF_EXPORT int cef_execute_process(const struct _cef_main_args_t* args,
    cef_app_t* application, void* windows_sandbox_info);

///
// This function should be called on the main application thread to initialize
// the CEF browser process. The |application| parameter may be NULL. A return
// value of true (1) indicates that it succeeded and false (0) indicates that it
// failed. The |windows_sandbox_info| parameter is only used on Windows and may
// be NULL (see cef_sandbox_win.h for details).
///
CEF_EXPORT int cef_initialize(const struct _cef_main_args_t* args,
    const struct _cef_settings_t* settings, cef_app_t* application,
    void* windows_sandbox_info);

///
// This function should be called on the main application thread to shut down
// the CEF browser process before the application exits.
///
CEF_EXPORT void cef_shutdown();

///
// Perform a single iteration of CEF message loop processing. This function is
// used to integrate the CEF message loop into an existing application message
// loop. Care must be taken to balance performance against excessive CPU usage.
// This function should only be called on the main application thread and only
// if cef_initialize() is called with a CefSettings.multi_threaded_message_loop
// value of false (0). This function will not block.
///
CEF_EXPORT void cef_do_message_loop_work();

///
// Run the CEF message loop. Use this function instead of an application-
// provided message loop to get the best balance between performance and CPU
// usage. This function should only be called on the main application thread and
// only if cef_initialize() is called with a
// CefSettings.multi_threaded_message_loop value of false (0). This function
// will block until a quit message is received by the system.
///
CEF_EXPORT void cef_run_message_loop();

///
// Quit the CEF message loop that was started by calling cef_run_message_loop().
// This function should only be called on the main application thread and only
// if cef_run_message_loop() was used.
///
CEF_EXPORT void cef_quit_message_loop();

///
// Set to true (1) before calling Windows APIs like TrackPopupMenu that enter a
// modal message loop. Set to false (0) after exiting the modal message loop.
///
CEF_EXPORT void cef_set_osmodal_loop(int osModalLoop);