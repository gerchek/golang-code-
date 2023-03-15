package constructor

import (
	"project/internal/domain/xml/controller"
	"project/internal/domain/xml/service"
	"project/internal/domain/xml/storage"

	"gorm.io/gorm"
)

var (
	XmlRepository storage.XmlStorage
	XmlService    service.XmlService
	XmlController controller.XmlController
)

func XmlRequirementsCreator(client *gorm.DB) {
	XmlRepository = storage.NewXmlStorage(client)
	XmlService = service.NewXmlService(XmlRepository)
	XmlController = controller.NewXmlController(XmlService)
}
