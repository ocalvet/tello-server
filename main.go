package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ocalvet/tello-testing/app"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/api"
	"gobot.io/x/gobot/platforms/dji/tello"
)

func main() {
	drone := tello.NewDriver("8888")
	master := gobot.NewMaster()
	a := api.NewAPI(master)
	a.Start()
	r := gin.Default()
	work := app.CreateWorker(drone, r)
	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		work,
	)
	master.AddRobot(robot)
	master.Start()
}
