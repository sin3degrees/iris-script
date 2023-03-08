package route

import (
	"net/http"

	"iris-script/conf"
	"iris-script/controllers"
	"iris-script/middleware"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func InitRouter(app *iris.Application) {
	baseUrl := "/api"
	mvc.New(app.Party(baseUrl + "/user")).Handle(controllers.NewUserController())
	if conf.Sysconfig.Common.AuthCheck {
		app.Use(middleware.GetJWT().Serve) // jwt
	}
	mvc.New(app.Party(baseUrl + "/book")).Handle(controllers.NewBookController())
	mvc.New(app.Party(baseUrl + "/script")).Handle(controllers.NewScriptController())
}

func CrossAccess11(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
func CrossAccess(ctx iris.Context) {
	ctx.ResponseWriter().Header().Set("Access-Control-Allow-Origin", "*")
}
