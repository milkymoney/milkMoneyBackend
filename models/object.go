package models

import (
	"errors"
	"strconv"
	"time"
	"fmt"
)

var (
	Objects map[string]*Object
)

type Object struct {
	ObjectId   string `orm:"pk"`
	Score      int64
	PlayerName string
}

func init() {
	Objects = make(map[string]*Object)
}

func AddOne(object Object) (ObjectId string) {
	object.ObjectId = "astaxie" + strconv.FormatInt(time.Now().UnixNano(), 10)
	o := getOrm()
	id, err := o.Insert(&object)
	fmt.Printf("Insert object with ID:%d, ERR:%v\n",id,err)
	Objects[object.ObjectId] = &object
	return object.ObjectId
}

func GetOne(ObjectId string) (object *Object, err error) {
	if v, ok := Objects[ObjectId]; ok {
		return v, nil
	}
	return nil, errors.New("ObjectId Not Exist")
}

func GetAll() map[string]*Object {
	return Objects
}

func Update(ObjectId string, Score int64) (err error) {
	if v, ok := Objects[ObjectId]; ok {
		v.Score = Score
		return nil
	}
	return errors.New("ObjectId Not Exist")
}

func Delete(ObjectId string) {
	delete(Objects, ObjectId)
}

