package services

import (
	"backend/clients"
	"backend/dao"
)

func GetCourseByID(id string) (dao.Course, error) {
	var course dao.Course
	result := clients.DB.Where("id = ?", id).First(&course)
	return course, result.Error
}

func GetCoursesByName(name string) ([]dao.Course, error) {
	var courses []dao.Course
	result := clients.DB.Where("nombre LIKE ?", "%"+name+"%").Find(&courses)
	return courses, result.Error
}
