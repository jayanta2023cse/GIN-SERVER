// Package middlewares handles application configuration, including loading environment variables.
package middlewares

import (
	"app/config"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var limiterStore = make(map[string]*rate.Limiter)
var mu sync.Mutex

// getLimiter returns a rate limiter for a specific IP
func getLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	if limiter, exists := limiterStore[ip]; exists {
		return limiter
	}

	requests, err := strconv.Atoi(config.AppConfig.ThrottleLimit)
	if err != nil {
		requests = 10
	}
	durationMs, err := strconv.Atoi(config.AppConfig.ThrottleTTL)
	if err != nil {
		durationMs = 60000
	}

	// Calculate rate: X requests per duration
	duration := time.Duration(durationMs) * time.Millisecond
	limiter := rate.NewLimiter(rate.Every(duration/time.Duration(requests)), requests)

	// Store for this IP
	limiterStore[ip] = limiter
	return limiter
}

func ThrottleGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := getLimiter(ip)

		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"status":  false,
				"error":   "Rate limit exceeded",
				"message": "Too many requests, please try again later.",
				"data":    map[string]interface{}{},
			})
			return
		}
		c.Next()
	}
}
