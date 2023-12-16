package repository

import (
	"crud-go/data/request"
	"crud-go/helper"
	"crud-go/model"
	"errors"

	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	DB *gorm.DB
}

func NewTagsRepositoryImpl(DB *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{DB: DB}
}

// Delete implements TagsRepository.
func (t *TagsRepositoryImpl) Delete(tagsId int) {
	var tags model.Tags
	result := t.DB.Where("id = ?", tagsId).Delete(&tags)
	helper.ErrorPanic(result.Error)
}

// FindAll implements TagsRepository.
func (t *TagsRepositoryImpl) FindAll() []model.Tags {
	var tags []model.Tags
	result := t.DB.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags
}

// FindById implements TagsRepository.
func (t *TagsRepositoryImpl) FindById(tagsId int) (tags model.Tags, err error) {
	var tag model.Tags
	result := t.DB.Find(&tag, tagsId)
	if result != nil {
		return tags, nil
	} else {
		return tags, errors.New("tag is not found")
	}

}

// Save implements TagsRepository.
func (t *TagsRepositoryImpl) Save(tags model.Tags) {
	result := t.DB.Create(&tags)
	helper.ErrorPanic(result.Error)
}

// Update implements TagsRepository.
func (t *TagsRepositoryImpl) Update(tags model.Tags) {
	var updateTag = request.UpdateTagsRequest{
		Id:   tags.Id,
		Name: tags.Name,
	}

	result := t.DB.Model(&tags).Updates(updateTag)
	helper.ErrorPanic(result.Error)
}
