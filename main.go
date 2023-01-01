package main

import (
	"flag"
	"log"
	"sendx/routes"
	"sendx/utils"
	"sendx/utils/cache"

	"github.com/gofiber/fiber/v2"

	"github.com/goccy/go-json"
)

func main() {

	var (
		maxQueueSize = flag.Int("max_queue_size", 100, "The size of job queue")
		maxWorkers   = flag.Int("max_workers", 5, "The number of workers to start")
	)
	flag.Parse()

	// create job channel
	jobs := make(chan utils.Job, *maxQueueSize)

	// create workers
	for i := 1; i <= *maxWorkers; i++ {
		go func(i int) {
			for j := range jobs {
				utils.DoWork(i, j, jobs)
			}
		}(i)
	}

	// Replace the default JSON encoder and decoder with the goccy/go-json encoder and decoder
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	cache.InitCache()

	// Register the routes
	routes.Register(app, jobs)

	log.Fatal(app.Listen(":7771"))
}
