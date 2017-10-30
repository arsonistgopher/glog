package main

import (
	"os"

	// This is the glog logging decoupling package.
	"github.com/arsonistgopher/glog"

	// This is the logging library I want to use for the example project.
	coreoslog "github.com/coreos/go-log/log"

	// Test package to show usage. Nothing important.
	useslogging "github.com/arsonistgopher/glog/examples/useslogging"
)

func main() {
	// Create new `logger` from glog and give it a name.
	logger := glog.Logger{Name: "GlogDemo"}

	// This logger is created using the coreos go-log New() function.
	// The type returned by coreoslog.New implements the interface in projectlogger.
	// This could be any instance of any other concrete type in other logging packages
	// so long as the 4x methods (Debug/Info/Critical/Error) are implemented.
	logger.LoggingBase = coreoslog.New(logger.Name, true, coreoslog.CombinedSink(os.Stdout, coreoslog.BasicFormat, coreoslog.BasicFields))

	// Now I can use Debug/Info/Critical/Error from the project logger,
	logger.Debug("Thing1")
	logger.Info("Thing2")

	// Call the function, send in its parameters and the logger type.
	// The logging package does not have any global vars, so you must pass back in the logger instance.
	useslogging.DoSomethingF1("Thing3", logger)

	// Another thing might be to create an variable of a concrete type with logging embedded and then just call funcs.
	t1 := useslogging.DoSomethingT{Name: "Thing4", Logger: logger}
	t1.DoSomethingF2("Thing5")

	// Note, no package has global vars. Beautiful!

}
