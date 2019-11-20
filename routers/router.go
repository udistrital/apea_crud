// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/simarhj/apea_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/observacion",
			beego.NSInclude(
				&controllers.ObservacionController{},
			),
		),

		beego.NSNamespace("/alerta",
			beego.NSInclude(
				&controllers.AlertaController{},
			),
		),

		beego.NSNamespace("/caso",
			beego.NSInclude(
				&controllers.CasoController{},
			),
		),

		beego.NSNamespace("/pregunta_opcion",
			beego.NSInclude(
				&controllers.PreguntaOpcionController{},
			),
		),

		beego.NSNamespace("/opcion",
			beego.NSInclude(
				&controllers.OpcionController{},
			),
		),

		beego.NSNamespace("/seguimiento",
			beego.NSInclude(
				&controllers.SeguimientoController{},
			),
		),

		beego.NSNamespace("/pregunta",
			beego.NSInclude(
				&controllers.PreguntaController{},
			),
		),

		beego.NSNamespace("/respuesta",
			beego.NSInclude(
				&controllers.RespuestaController{},
			),
		),

		beego.NSNamespace("/estudiante",
			beego.NSInclude(
				&controllers.EstudianteController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
