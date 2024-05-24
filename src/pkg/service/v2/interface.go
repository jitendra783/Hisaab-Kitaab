package v2

import "hisaab-kitaab/pkg/db"

type ServiceObj struct {
}

func NewServiceGroup(db db.DBLayer) ServiceLayer{ 
	return nil
}

type ServiceLayer interface{}