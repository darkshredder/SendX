package handlers

import (
	"fmt"
	"math"
	"sendx/utils"
	"sendx/utils/cache"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PageSourcePayload struct {
	URI        string `json:"uri"`
	RetryLimit int    `json:"retryLimit"` // optional
}

// GetPageSource is the handler for the POST /pagesource route
func GetPageSource(c *fiber.Ctx, jobs chan utils.Job) error {

	payload := new(PageSourcePayload)

	if payload.RetryLimit == 0 {
		payload.RetryLimit = 10
	}

	payload.RetryLimit = int(math.Min(float64(payload.RetryLimit), float64(10)))

	if err := c.BodyParser(payload); err != nil {
		return err
	}

	uuid := uuid.New()

	cacheVal := cache.MyCache.GetCacheItem(payload.URI)
	if cacheVal != nil {
		return c.JSON(fiber.Map{
			"uri":       payload.URI,
			"id":        cacheVal.Data().(string),
			"sourceUri": "/files/" + cacheVal.Data().(string) + ".html",
		})
	}

	cache.MyCache.AddCacheItem(payload.URI, uuid.String())

	job := utils.Job{Name: "GetPageSource", Retries: payload.RetryLimit, Functions: func() bool {
		return utils.DownloadPageSourceJob(payload.URI, uuid.String())
	}, Tries: 1,
	}

	go func() {
		fmt.Printf("added: %s\n", job.Name)
		jobs <- job
	}()

	return c.JSON(fiber.Map{
		"uri":       payload.URI,
		"id":        uuid.String(),
		"sourceUri": "/files/" + uuid.String() + ".html",
	})

}
