package app

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gobot.io/x/gobot/platforms/dji/tello"
)

func CreateWorker(drone *tello.Driver, r *gin.Engine) func() {
	return func() {
		drone.TakeOff()
		r.Use(cors.New(cors.Config{
			AllowAllOrigins:  true,
			AllowCredentials: true,
			AllowMethods:     []string{"*"},
			AllowHeaders:     []string{"*"},
		}))
		r.GET("/tello/takeoff", func(c *gin.Context) {
			drone.TakeOff()
			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		})

		r.GET("/tello/land", func(c *gin.Context) {
			drone.Land()
			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		})

		r.GET("/tello/left", func(c *gin.Context) {
			drone.Left(10)
			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		})

		r.GET("/tello/right", func(c *gin.Context) {
			drone.Right(10)
			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		})

		r.GET("/tello/forward", func(c *gin.Context) {
			drone.Forward(10)
			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		})

		r.GET("/tello/backward", func(c *gin.Context) {
			drone.Backward(10)
			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		})

		r.GET("/tello/flip/right", func(c *gin.Context) {
			drone.RightFlip()
			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		})

		r.GET("/tello/flip/left", func(c *gin.Context) {
			drone.LeftFlip()
			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		})

		r.GET("/tello/flip/front", func(c *gin.Context) {
			drone.FrontFlip()
			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		})

		r.GET("/tello/flip/back", func(c *gin.Context) {
			drone.BackFlip()
			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		})

		r.GET("/tello/startfeed", func(c *gin.Context) {
			// drone.StartVideo()
			// drone.SetVideoEncoderRate(4)
			// gobot.Every(100*time.Millisecond, func() {
			// 				drone.StartVideo()
			// })
			drone.FrontFlip()
			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		})

		r.GET("/tello/endfeed", func(c *gin.Context) {
			drone.BackFlip()
			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		})

		r.Run(":9081")
	}
}
