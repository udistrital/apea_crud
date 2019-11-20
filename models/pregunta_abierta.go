package models

type PreguntaAbierta struct {
	Id_RENAME int16  `orm:"column(id)"`
	Pregunta  string `orm:"column(pregunta);null"`
	IdOpcion  int16  `orm:"column(id_opcion);null"`
}
