package main

import "github.com/manish-neemnarayan/go-native-server/types"

type DBHouse struct {
	data map[int]string
}

func NewDBHouse() *DBHouse {
	return &DBHouse{
		data: make(map[int]string),
	}
}

func (d *DBHouse) insert(data *types.PostData) error {
	d.data[data.Id] = data.Val1
	return nil
}

func (d *DBHouse) get(id int) (resData string, err error) {
	resData, ok := d.data[id]
	if !ok {
		return "", err
	}
	return resData, nil
}
