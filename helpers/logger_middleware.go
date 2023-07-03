package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

type LogData struct {
	Method       string        `json:"method"`
	URI          string        `json:"uri"`
	IP           string        `json:"ip"`
	Headers      http.Header   `json:"headers"`
	QueryParams  url.Values    `json:"queryParams"`
	RequestBody  string        `json:"requestBody"`
	ResponseCode int           `json:"responseCode"`
	ResponseBody string        `json:"responseBody"`
	Latency      time.Duration `json:"latency"`
}

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func RequestResponseLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start time
		writer := &responseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = writer

		// Start time
		start := time.Now()

		// Get the request IP address
		ip := c.ClientIP()

		// Read the request body
		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body.Close()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))

		c.Next()

		// Get the response body
		resBody := writer.body.String()

		// Calculate the response time
		latency := time.Since(start)

		logData := LogData{
			Method:       c.Request.Method,
			URI:          c.Request.RequestURI,
			IP:           ip,
			Headers:      c.Request.Header,
			QueryParams:  c.Request.URL.Query(),
			RequestBody:  string(reqBody),
			ResponseBody: resBody,
			ResponseCode: c.Writer.Status(),
			Latency:      latency,
		}

		logJSON, _ := json.Marshal(logData)
		fmt.Printf("Response: %s\n", logJSON)
	}
}
