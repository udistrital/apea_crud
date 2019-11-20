package models

type RespuestaAbierta struct {
	Id_RENAME        int16  `orm:"column(id)"`
	CodigoEstudiante int64  `orm:"column(codigo_estudiante);null"`
	IdPregunta       int16  `orm:"column(id_pregunta);null"`
	Respuesta        string `orm:"column(respuesta);null"`
}
