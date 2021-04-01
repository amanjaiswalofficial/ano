package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

// API is a custom type over *gin.Engine to add new methods on the existing type
type API struct{ Obj *gin.Engine }

// InitAPI will initialize an API variable and return the same
func InitAPI() API {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowHeaders:     []string{"Origin"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        AllowOriginFunc: func(origin string) bool {
            return origin == "http://localhost:3001"
        },
    }))

	// Registering methods to APIs
	r.GET("/feed", FeedData)

	return API{r}
}

// Start will take an API variable as receiver and start running it as server
func (a *API) Start(port int) {
	a.Obj.Run(":" + strconv.Itoa(port))
}

// FeedData will return a JSON of all the RSS data stored in the source file
func FeedData(c *gin.Context) {
	result := readDataFromFile("data.json")
	if result != nil {
		c.JSON(200, gin.H{"response": result, "error": false})
	} else {
		c.JSON(404, gin.H{"response": nil, "error": true})
	}

}
