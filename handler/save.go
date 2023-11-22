package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
)

const urlPosts = "https://jsonplaceholder.typicode.com/posts"

type post struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	UsrID int    `json:"userId"`
}

func HandlePostsSave(c *fiber.Ctx) error {
	postData := c.Body()
	var p post
	json.Unmarshal(postData, &p)
	if p.UsrID == 0 || len(p.Body) == 0 || len(p.Title) == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("Ill-formated post")
	}
	postData, _ = json.Marshal(p)
	body := bytes.NewReader(postData)

	req, _ := http.NewRequest(http.MethodPost, urlPosts, body)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	defer resp.Body.Close()
	//io.Copy(io.Discard, resp.Body)
	responseData, _ := io.ReadAll(resp.Body)
	return c.SendString(string(responseData))
}
