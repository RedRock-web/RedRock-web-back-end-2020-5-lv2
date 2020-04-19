package router

import (
	"RedRock-web-back-end-2020-5-lv2/database"
	"errors"
	"github.com/gin-gonic/gin"
)

type GetForm struct {
	types string
	data  string
}

func Handle(c *gin.Context) {
	g := GetForm{}
	if err := c.ShouldBindJSON(&g); err != nil {
		errors.New("bind json error!")
		FormError(c)
		return
	}
	var student database.Student
	var classs []database.Class

	if g.types == "int" {
		database.G_db.Where("student_id = ?", g.data).Find(&student)
		database.G_db.Where("student_id = ?", g.data).Find(&classs)
	} else if g.types == "string" {
		database.G_db.Where("student_name = ?", g.data).Find(&student)
		database.G_db.First("student_name = ?", g.data).Find(&classs)
	} else {
		FormError(c)
	}

	var data []gin.H
	var temp []gin.H

	for _, class := range classs {
		temp = append(temp, gin.H{
			"lesson":   class.Lesson,
			"rasWeek":  class.RawWeek,
			"location": class.Location,
			"teacher":  class.Teacher,
			"day":      class.Day,
		})
	}

	data = append(data, gin.H{
		"studentId":   student.StudentId,
		"studentName": student.StudentName,
		"day":         student.Day,
		"semester":    student.Semester,
		"class":       temp,
	})
	OkWithData(c, data)
}
