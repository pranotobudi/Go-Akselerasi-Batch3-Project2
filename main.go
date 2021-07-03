package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/routes"
)

func main() {

	fmt.Println("bismillah")
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("failed to load .env file")
	}
	// cron := gocron.NewScheduler(time.Now().Location())
	// cron.Every(5).Seconds().Do(func() { fmt.Println("alhamdulillah") })
	// fmt.Println("=========MAIN ROUTE CRON bg : %+v \n", cron)
	// // starts the scheduler and blocks current execution path
	// cron.StartAsync()

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Pre(middleware.RemoveTrailingSlash())
	routes.DefineApiRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))

}
func JustSchedule() {
	fmt.Println("Bismillah")
}
