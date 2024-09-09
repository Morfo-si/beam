package server

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/morfo-si/beam/internal/config"
	"github.com/stretchr/testify/assert"
)

func Test_Hello(t *testing.T) {
	hello := "hello"
	assert.Equal(t, hello, "hello")
}

func TestFiberApp(t *testing.T) {
	app := fiber.New()

	app.Get("/test", func(ctx fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})

	r := httptest.NewRequest("GET", "/test", nil)
	resp, _ := app.Test(r, -1)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestACEServer(t *testing.T) {
	srv := NewACEServer()
	app := srv.App()

	var chatRequest config.ChatRequest
	chatRequest.Prompt = "a prompt"
	chatRequest.Question = "a question"
	content, _ := json.Marshal(&chatRequest)

	req := httptest.NewRequest("POST", "/api/v1/chat", bytes.NewReader(content))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.FormatInt(req.ContentLength, 10))
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)

	var chatResponse config.OllamaResponse
	json.Unmarshal(body, &chatResponse)
	assert.True(t, len(chatResponse.Response) > 0)
}
