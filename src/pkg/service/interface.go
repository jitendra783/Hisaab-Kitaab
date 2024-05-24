package service

import (
	"hisaab-kitaab/pkg/db"
	v1 "hisaab-kitaab/pkg/service/v1"
	v2 "hisaab-kitaab/pkg/service/v2"
)

type serviceObj struct {
	V1 v1.ServiceLayer
	V2 v2.ServiceLayer
}

func NewServiceGroupObject(db db.DBLayer) ServiceGroupLayer {
	temp := &serviceObj{v1.NewServiceGroup(db),
		v2.NewServiceGroup(db),
	}
	return temp
}

type ServiceGroupLayer interface {
	GetV1Service() v1.ServiceLayer
	GetV2Service() v2.ServiceLayer
}

func (s *serviceObj) GetV1Service() v1.ServiceLayer {
	return s.V1
}

func (s *serviceObj) GetV2Service() v2.ServiceLayer {
	return s.V2
}
