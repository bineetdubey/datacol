package client

import (
	"github.com/dinesh/datacol/client/models"
	"time"
)

var (
	b_bucket = []byte("builds")
)

func (c *Client) NewBuild(app *models.App) *models.Build {
	b := &models.Build{
		App:       app.Name,
		Id:        generateId("B", 5),
		Status:    "creating",
		CreatedAt: time.Now(),
	}

	return b
}

func (c *Client) GetBuilds(app string) (models.Builds, error) {
	return c.Provider().BuildList(app, 20)
}

func (c *Client) GetBuild(app, id string) (*models.Build, error) {
	return c.Provider().BuildGet(app, id)
}

func (c *Client) DeleteBuild(app, id string) error {
	return c.Provider().BuildDelete(app, id)
}
