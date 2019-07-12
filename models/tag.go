package models

import (
	"fmt"
)

type Tag struct {
	Model
	Name 		string 	`json:"name" binding:"required"`
	CreatedBy 	string 	`json:"create_by"`
	ModifiedBy 	string 	`json:"modified_by"`
	State 		int 	`json:"state binding:"required,stateEnum"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagsTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExitTagName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(tag Tag) bool {
	db.Create(&tag)
	return true
}