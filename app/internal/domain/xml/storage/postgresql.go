package storage

import (
	"project/internal/model"

	"gorm.io/gorm"
)

type XmlStorage interface {
	All() []*model.Xml
	CreateOrUpdate(xml *model.Xml, id int) error
}

type xmlStorage struct {
	client *gorm.DB
}

func NewXmlStorage(client *gorm.DB) XmlStorage {
	return &xmlStorage{
		client: client,
	}
}

func (ps *xmlStorage) All() []*model.Xml {
	var xml []*model.Xml
	ps.client.Find(&xml)
	return xml
}

func (ps *xmlStorage) CreateOrUpdate(xml *model.Xml, id int) error {
	existingUser := &model.Xml{}
	ps.client.Where("uid = ?", id).First(existingUser)
	if existingUser.ID != 0 {
		err1 := ps.client.Model(existingUser).Updates(xml)
		if err1.Error != nil {
			return err1.Error
		}
	} else {
		err2 := ps.client.Create(xml)
		if err2.Error != nil {
			return err2.Error
		}
	}
	return nil

}
