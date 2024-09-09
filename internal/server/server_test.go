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

func BuildChatRequest(prompt string, question string) ([]byte, error) {
	var chatRequest config.ChatRequest
	chatRequest.Prompt = prompt
	chatRequest.Question = question
	return json.Marshal(&chatRequest)
}

func SendChatRequest(app *fiber.App, content []byte) (*http.Response, error) {
	req := httptest.NewRequest("POST", "/api/v1/chat", bytes.NewReader(content))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.FormatInt(req.ContentLength, 10))
	return app.Test(req, -1)
}

func ParseChatResponse(resp *http.Response, chatResponse *config.OllamaResponse) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, chatResponse)
}

func SendOLSRequest(app *fiber.App, content []byte) (*http.Response, error) {
	req := httptest.NewRequest("POST", "/api/v1/ols", bytes.NewReader(content))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.FormatInt(req.ContentLength, 10))
	return app.Test(req, -1)
}

func ParseOLSResponse(resp *http.Response, olsResponse *config.OLSResponse) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, olsResponse)
}
func TestACEServer(t *testing.T) {
	srv := NewACEServer()
	app := srv.App()

	content, _ := BuildChatRequest("a prompt", "a question")
	resp, _ := SendChatRequest(app, content)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var chatResponse config.OllamaResponse
	ParseChatResponse(resp, &chatResponse)
	assert.True(t, len(chatResponse.Response) > 0)
}

func TestACEServerWithoutPrompt(t *testing.T) {
	srv := NewACEServer()
	app := srv.App()

	content, _ := BuildChatRequest("", "a question")
	resp, _ := SendChatRequest(app, content)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestACEServerWithoutQuestion(t *testing.T) {
	srv := NewACEServer()
	app := srv.App()

	content, _ := BuildChatRequest("a prompt", "")
	resp, _ := SendChatRequest(app, content)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestACEServerWithOLS(t *testing.T) {
	srv := NewACEServer()
	app := srv.App()

	content, _ := BuildChatRequest("", "What is Ansible?")
	resp, _ := SendOLSRequest(app, content)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var olsResponse config.OLSResponse
	ParseOLSResponse(resp, &olsResponse)
	assert.True(t, len(olsResponse.Response) > 0)
}
