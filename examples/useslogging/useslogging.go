package useslogging

import (
	projectlogger "ghost2loggerservice/ghost2loggerd/scratches/logdepinject/logging"
)

// DoSomethingF1 returns a doubled up string.
func DoSomethingF1(message string, l projectlogger.Logger) string {
	l.Debug(message)
	return message
}

// DoSomethingT is a type with an embedded logger.
type DoSomethingT struct {
	projectlogger.Logger
	Name string
}

// DoSomethingF2 has an input message and pushes out to a logger.
func (d DoSomethingT) DoSomethingF2(message string) {
	d.Debug(d.Name)
	d.Debug(message)
}