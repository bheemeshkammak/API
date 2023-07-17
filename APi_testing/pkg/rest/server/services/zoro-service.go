package services

import (
	"github.com/bheemeshkammak/API/api_testing/pkg/rest/server/daos"
	"github.com/bheemeshkammak/API/api_testing/pkg/rest/server/models"
)

type ZoroService struct {
	zoroDao *daos.ZoroDao
}

func NewZoroService() (*ZoroService, error) {
	zoroDao, err := daos.NewZoroDao()
	if err != nil {
		return nil, err
	}
	return &ZoroService{
		zoroDao: zoroDao,
	}, nil
}

func (zoroService *ZoroService) CreateZoro(zoro *models.Zoro) (*models.Zoro, error) {
	return zoroService.zoroDao.CreateZoro(zoro)
}

func (zoroService *ZoroService) UpdateZoro(id int64, zoro *models.Zoro) (*models.Zoro, error) {
	return zoroService.zoroDao.UpdateZoro(id, zoro)
}

func (zoroService *ZoroService) DeleteZoro(id int64) error {
	return zoroService.zoroDao.DeleteZoro(id)
}

func (zoroService *ZoroService) ListZoros() ([]*models.Zoro, error) {
	return zoroService.zoroDao.ListZoros()
}

func (zoroService *ZoroService) GetZoro(id int64) (*models.Zoro, error) {
	return zoroService.zoroDao.GetZoro(id)
}
