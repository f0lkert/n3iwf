package app

import (
	n3iwf_context "github.com/f0lkert/n3iwf/internal/context"
	"github.com/f0lkert/n3iwf/pkg/factory"
)

type App interface {
	SetLogEnable(enable bool)
	SetLogLevel(level string)
	SetReportCaller(reportCaller bool)

	Start()
	Terminate()

	Context() *n3iwf_context.N3IWFContext
	Config() *factory.Config
}
