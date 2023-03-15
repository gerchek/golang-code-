package constructor

import (
	xmlConstructor "project/internal/domain/xml/constructor"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SetConstructor(client *gorm.DB, logger *logrus.Logger) {
	xmlConstructor.XmlRequirementsCreator(client)
}
