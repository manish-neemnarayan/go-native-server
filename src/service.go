package main

import (
	"github.com/manish-neemnarayan/go-native-server/types"
)

type DBHouser interface {
	insert(resData *types.PostData) error
	get(Id int) (string, error)
}

type Service struct {
	db DBHouser
}

func NewService() *Service {
	db := NewDBHouse()
	return &Service{
		db: db,
	}
}

func (s *Service) Insert(data *types.PostData) error {
	if err := s.db.insert(data); err != nil {
		return err
	}

	return nil
}

func (s *Service) Get(id int) (res *types.PostData, err error) {
	val, err := s.db.get(id)
	if err != nil {
		return &types.PostData{}, err
	}

	return &types.PostData{
		Id:   id,
		Val1: val,
	}, nil
}
