package dto

type XmlDTO struct {
	Name string `json:"name" form:"name" validate:"required,max=70"`
}
