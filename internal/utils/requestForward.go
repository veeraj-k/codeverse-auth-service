package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func ForwardRequest(c *gin.Context, targetURL string) {
	targetURL = targetURL + c.Request.URL.Path

	req, err := http.NewRequest(c.Request.Method, targetURL, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	// Copy headers from the original request
	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	fmt.Println("Forwarding request to:", targetURL)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to forward request"})
		return
	}
	defer resp.Body.Close()

	// Copy response headers
	for key, values := range resp.Header {
		for _, value := range values {
			c.Writer.Header().Add(key, value)
		}
	}

	// Copy response status code
	c.Writer.WriteHeader(resp.StatusCode)

	// Copy response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}
	c.Writer.Write(body)
}

func ForwardRequestV2(c *gin.Context, target string) {
	targetURL, err := url.Parse(target)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid service URL"})
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	fmt.Println(c.Request.URL.RawPath)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header.Clone()
		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
		req.URL.Path = c.Request.URL.Path
		req.Host = targetURL.Host
	}
	fmt.Println("Forwarding request to:", targetURL)

	proxy.ServeHTTP(c.Writer, c.Request)

}
