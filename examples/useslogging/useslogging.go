package useslogging

import (
	"github.com/arsonistgopher/glog"
)

// DoSomethingF1 returns a string.
func DoSomethingF1(message string, l glog.Logger) string {
	l.Debug(message)
	return message
}

// DoSomethingT is a type with an embedded logger.
type DoSomethingT struct {
	glog.Logger
	Name string
}

// DoSomethingF2 has an input message and pushes out to a logger.
func (d DoSomethingT) DoSomethingF2(message string) {
	d.Debug(d.Name)
	d.Debug(message)
}
