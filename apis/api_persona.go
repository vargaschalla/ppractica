package apis

import (
	"net/http"

	"PRACTICA ALEX/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PersonaGetId(c *gin.Context) {
	id := c.Params.ByName("id")
	var cur models.Persona
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	if err := conn.First(&cur, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &cur)
}

func PersonaIndex(c *gin.Context) {
	var lis []models.Persona
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Find(&lis)
	c.JSON(http.StatusOK, lis)
}
func PersonaPost(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	var per models.Persona
	//est := models.Estudiante{Name: c.PostForm("name"),Paternal: c.PostForm("paternal"),Maternal: c.PostForm("maternal"), Age: c.PostForm("age"), State: c.PostForm("state"),}
	if err := c.BindJSON(&est); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Create(&per)
	c.JSON(http.StatusOK, &per)
}

func PersonaPut(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var per models.Persona
	if err := conn.First(&p, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	per.Nombre = c.PostForm("nombre")
	per.Paterno = c.PostForm("paterno")
	per.Materno = c.PostForm("materno")
	per.Edad = c.PostForm("edad")
	per.Fechanacimiento = c.PostForm("fechanacimiento")
	per.Estadocivil = c.PostForm("estadocivil")
	c.BindJSON(&per)
	conn.Save(&per)
	c.JSON(http.StatusOK, &per)
}

func PersonaDelete(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var p models.Persona
	if err := conn.Where("id = ?", id).First(&p).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Unscoped().Delete(&p)
}
