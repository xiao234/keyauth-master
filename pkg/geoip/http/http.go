package http

import (
	"errors"

	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/geoip"
)

var (
	api = &handler{}
)

type handler struct {
	service geoip.Service
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	geoipRouter := router.ResourceRouter("IP")
	geoipRouter.BasePath("geoip")
	geoipRouter.Handle("GET", "/query", h.LoopupIP).AddLabel(label.Get)

	geoipRouter.Permission(true)
	geoipRouter.Handle("POST", "/dbfile", h.UpdateDBFile).AddLabel(label.Create)

}

func (h *handler) Config() error {
	if pkg.Department == nil {
		return errors.New("denpence department service is nil")
	}

	h.service = pkg.GEOIP
	return nil
}

func init() {
	pkg.RegistryHTTPV1("geoip", api)
}
