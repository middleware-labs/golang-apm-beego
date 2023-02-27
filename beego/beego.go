package beego

import (
	"github.com/astaxie/beego"
	"github.com/middleware-labs/golang-apm/tracker"
	"go.opentelemetry.io/contrib/instrumentation/github.com/astaxie/beego/otelbeego"
)

func Middleware(config *tracker.Config) beego.MiddleWare {
	return otelbeego.NewOTelBeegoMiddleWare(config.ServiceName)
}
