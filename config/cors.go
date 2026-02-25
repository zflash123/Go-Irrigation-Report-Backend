package config

import (
	"github.com/rs/cors"
)

var CorsObject = cors.New(cors.Options{
	AllowedOrigins: []string{"http://irrigation-report.vercel.app", "https://irrigation-report.vercel.app"},
	// AllowedHeaders: []string{"authorization"},
	AllowedHeaders: []string{"*"},
	AllowCredentials: true,
	Debug: false,
})