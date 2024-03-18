package main

import (
	"errors"
	"net/http"

	"github.com/Joanoni/inutil"
)

type ExecuteInput struct {
	Method  string   `json:"method"`
	Url     string   `json:"url"`
	Header  []Header `json:"header"`
	Payload any      `json:"payload"`
}

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ExecuteOutput struct {
	Body       any `json:"body"`
	StatusCode int `json:"statusCode"`
}

func main() {
	s := inutil.Start(&inutil.StartInput{
		Server: &inutil.StartServerInput{
			Port: ":7963",
		},
		Log: &inutil.StartLogInput{
			InternalLog: inutil.StartLogEnvInput{
				Development: true,
				Stage:       true,
				Production:  true,
			},
			DebugLog: inutil.StartLogEnvInput{
				Development: true,
				Stage:       true,
				Production:  true,
			},
			TimeFormat: inutil.LogFormat,
		},
		Enviroment: inutil.Enviroment_Development,
	})

	s.Server.Use(inutil.MiddlewareCors())
	s.Server.Use(inutil.MiddlewareLog())
	s.Server.Use(inutil.MiddlewareSafety())

	s.Server.Get("/api/test/get", okHandler)
	s.Server.Get("/api/test/error", errorHandler)
	s.Server.Post("/api/test/post", postHandler)
	s.Server.Post("/api/execute", executeHandler)
	s.Server.Run()
}

func okHandler(c *inutil.Context) {
	data := map[string]any{
		"hello": "world",
	}
	inutil.JSON(inutil.Return[map[string]any]{
		Message:    "success",
		Data:       &data,
		Success:    true,
		StatusCode: inutil.StatusOK,
	}, c)
}

func postHandler(c *inutil.Context) {
	var body any

	err := c.Body(&body)
	if c.HandleError(err) {
		return
	}

	inutil.LogDebug(body)

	inutil.JSON(inutil.Return[any]{
		Message:    "success",
		Data:       &body,
		Success:    true,
		StatusCode: inutil.StatusOK,
	}, c)
}

func errorHandler(c *inutil.Context) {
	c.HandleError(errors.New("erro feio"))
}

func executeHandler(c *inutil.Context) {
	var input ExecuteInput

	err := c.Body(&input)
	if c.HandleError(err) {
		return
	}

	inutil.LogDebug(input)

	headers := http.Header{}

	for _, header := range input.Header {
		inutil.LogDebug(header)
		headers.Set(header.Name, header.Value)
	}

	var resp inutil.Return[inutil.RequestReponse[*any]]

	if input.Payload == nil {
		resp = inutil.Request[any](inutil.RequestInput{
			Method:  input.Method,
			Url:     input.Url,
			Payload: nil,
			Header:  headers,
		}, c)
	} else {
		resp = inutil.Request[any](inutil.RequestInput{
			Method: input.Method,
			Url:    input.Url,
			Payload: &inutil.RequestPayloadInput{
				Body:        input.Payload,
				ContentType: inutil.ApplicationJSON,
			},
			Header: headers,
		}, c)
	}

	inutil.LogDebug(resp)

	inutil.JSON(inutil.Return[ExecuteOutput]{
		Message: resp.Message,
		Data: &ExecuteOutput{
			Body:       resp.Data.Body,
			StatusCode: resp.Data.StatusCode,
		},
		Success:    true,
		StatusCode: inutil.StatusOK,
	}, c)
}
