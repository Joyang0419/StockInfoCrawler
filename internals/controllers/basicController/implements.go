package basicController

import (
	"StockInfoCrawler/internals/models"
	"StockInfoCrawler/internals/repos/basicRepo"
	"StockInfoCrawler/internals/services/basicService"
	log "github.com/sirupsen/logrus"
)

type ControllerBasic struct {
	Service basicService.InterfaceBasicService
	Repo    basicRepo.InterfaceBasicRepo
}

func NewControllerBasic(Service basicService.InterfaceBasicService, Repo basicRepo.InterfaceBasicRepo) (
	controller *ControllerBasic) {
	controller = &ControllerBasic{
		Service: Service,
		Repo:    Repo,
	}
	return controller
}

func (controller *ControllerBasic) ScrapeCategory() {
	categories, err := controller.Service.ScrapeCategory()
	if err != nil {
		log.Panic(err)
	}
	controller.Repo.CreateCategories(categories)
}

func (controller *ControllerBasic) ScrapeBasic() {
	categoryIDMapping := controller.Repo.GetCategoryIDMapping()
	basicModelChannel := make(chan models.BasicModel)
	err := controller.Service.ScrapeBasic(categoryIDMapping, basicModelChannel)
	if err != nil {
		log.Panic(err)
	}

	controller.Repo.CreateBasic(basicModelChannel)
}
