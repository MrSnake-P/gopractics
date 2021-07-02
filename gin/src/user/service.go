package user

import (
	"github.com/jinzhu/gorm"
	"github.com/linshenqi/sptty"
)

type Service struct {
	sptty.BaseService

	db      *gorm.DB

}

func (s *Service) Init(app sptty.ISptty) error {
	s.db = app.Model().(*sptty.ModelService).DB()


	app.AddRoute("GET", "/hello", s.helloworld)


	return nil
}