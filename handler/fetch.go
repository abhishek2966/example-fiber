package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	fiber "github.com/gofiber/fiber/v2"
)

const urlPhotos = "https://jsonplaceholder.typicode.com/photos"

type photo struct {
	AlbumID    int    `json:"albumId"`
	PhotoID    int    `json:"id"`
	PhotoTitle string `json:"title"`
	PhotoURL   string `json:"url"`
}

func HandlePhotosFetch(c *fiber.Ctx) error {
	albumIDString := c.Query("album_id")
	albumID, err := strconv.Atoi(albumIDString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("album_id required: %v", err))
	}
	resp, err := client.Get(urlPhotos)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("couldn't read response body: %v", err))
	}
	pics := []photo{}
	json.Unmarshal(data, &pics)
	picsFiltered := []photo{}
	for _, pic := range pics {
		if pic.AlbumID == albumID {
			picsFiltered = append(picsFiltered, pic)
		}
	}
	return c.JSON(picsFiltered)
}
