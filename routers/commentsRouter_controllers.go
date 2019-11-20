package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:AlertaController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:AlertaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:AlertaController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:AlertaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:AlertaController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:AlertaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:AlertaController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:AlertaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:AlertaController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:AlertaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:CasoController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:CasoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:CasoController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:CasoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:CasoController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:CasoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:CasoController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:CasoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:CasoController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:CasoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:EstudianteController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:EstudianteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:EstudianteController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:EstudianteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:EstudianteController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:EstudianteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:EstudianteController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:EstudianteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:EstudianteController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:EstudianteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:ObservacionController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:ObservacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:ObservacionController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:ObservacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:ObservacionController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:ObservacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:ObservacionController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:ObservacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:ObservacionController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:ObservacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:OpcionController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:OpcionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:OpcionController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:OpcionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:OpcionController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:OpcionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:OpcionController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:OpcionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:OpcionController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:OpcionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:PreguntaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:PreguntaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:PreguntaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:PreguntaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:PreguntaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:PreguntaOpcionController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:PreguntaOpcionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:PreguntaOpcionController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:PreguntaOpcionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:PreguntaOpcionController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:PreguntaOpcionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:PreguntaOpcionController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:PreguntaOpcionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:PreguntaOpcionController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:PreguntaOpcionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:RespuestaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:RespuestaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:RespuestaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:RespuestaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:RespuestaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:SeguimientoController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:SeguimientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:SeguimientoController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:SeguimientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:SeguimientoController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:SeguimientoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:SeguimientoController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:SeguimientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:SeguimientoController"] = append(beego.GlobalControllerRouter["github.com/simarhj/apea_crud/controllers:SeguimientoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
