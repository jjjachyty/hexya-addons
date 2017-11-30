package procurement

import (
	"github.com/hexya-erp/hexya/hexya/server"
	// Procurement dependencies
	_ "github.com/hexya-erp/hexya-addons/product"
)

const MODULE_NAME string = "procurement"

func init() {
	server.RegisterModule(&server.Module{
		Name:     MODULE_NAME,
		PostInit: func() {},
	})
}
