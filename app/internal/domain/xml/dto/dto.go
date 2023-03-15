package dto

type XmlDTO struct {
	Name string `json:"name" form:"name" validate:"required"`
	Type string `json:"type" form:"type" validate:"required"`
}
