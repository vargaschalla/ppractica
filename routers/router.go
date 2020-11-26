package routers

import (
	"crudjos/apis"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {

	conn, err := connectDB()
	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	r := gin.Default()

	r.Use(CORSMiddleware())

	r.Use(dbMiddleware(*conn))

	per := r.Group("/per")
	{
		per.GET("/persons/:id", apis.PersonaGetId)
		per.GET("/persons/", apis.PersonaIndex)
		per.POST("/persons/", apis.PersonaPost)
		per.PUT("/persons/:id", apis.PersonaPut)
		per.DELETE("/persons/:id", apis.PersonaDelete)
	}

	return r
}

func connectDB() (c *gorm.DB, err error) {

	dsn := "root:aracelybriguit@tcp(localhost:3306)/academico?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	conn.AutoMigrate(&models.Persona{})

	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	return conn, err
}

func dbMiddleware(conn gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.Header("Access-Control-Allow-Origin", "http://localhost","http://localhost:8080")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST. OPTIONS, GET, PUT, DELETE")

		if c.Request.Method {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
