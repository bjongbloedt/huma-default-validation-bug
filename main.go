package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"regexp"

	huma "github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
)

type Address struct {
	Line1       string `json:"line1" required:"true" minLength:"1" maxLength:"50"`
	Line2       string `json:"line2,omitempty" required:"false" minLength:"0" maxLength:"50" default:""`
	City        string `json:"city" required:"true" minLength:"1" maxLength:"64"`
	State       string `json:"state" required:"true" minLength:"1" maxLength:"32"`
	Zip         string `json:"zip" required:"true" minLength:"1" maxLength:"16"`
	CountryCode string `json:"countryCode" required:"false" minLength:"1" maxLength:"2" default:"US"`
}

func (a *Address) Resolve(_ huma.Context, prefix *huma.PathBuffer) []error {
	var errors []error
	if a.CountryCode == "US" {
		// Test US ZipCode formatting
		pattern := `^[0-9]{5}(?:-[0-9]{4})?$`
		match, _ := regexp.MatchString(pattern, a.Zip)
		if !match {
			errors = append(errors, &huma.ErrorDetail{
				Message:  "Zip has a bad value for country US",
				Location: prefix.With("zip"),
				Value:    a.Zip,
			})
		}

		// Test US state matching
		stateLen := len(a.State)
		if stateLen != 2 {
			errors = append(errors, &huma.ErrorDetail{
				Message:  "State has a bad value for country US (should be 2 characters)",
				Location: prefix.With("state"),
				Value:    a.State,
			})
		}

	}
	return errors
}

type TestRequest struct {
	Name        string   `json:"name"`
	Age         int      `json:"age"`
	HomeAddress *Address `json:"home" required:"true"`
	AwayAddress *Address `json:"away" required:"true"`
}

type TestInput struct {
	Body TestRequest `required:"true"`
}

func main() {
	// Create a new router & API
	router := chi.NewMux()
	api := humachi.New(router, huma.DefaultConfig("My API", "1.0.0"))

	huma.Post(api, "/test", func(ctx context.Context, i *TestInput) (*struct{}, error) {

		slog.Info("hey this worked", "input", i)

		return &struct{}{}, nil

	})

	// Start the server!
	http.ListenAndServe("127.0.0.1:8888", router)
}
