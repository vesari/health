package url

import (
	"net/http"

	"github.com/dimiro1/health"
)

// Checker is a checker that check a given URL
type Checker struct {
	URL string
}

// NewChecker returns a new url.Checker with the given URL
func NewChecker(url string) Checker {
	return Checker{URL: url}
}

// Check makes a HEAD request to the given URL
// If the request returns something different than StatusOK,
// The status is set to StatusBadRequest and the URL is considered Down
func (u Checker) Check() health.Health {
	req, err := http.NewRequest("HEAD", u.URL, nil)

	health := health.NewHealth()
	health.Up()

	resp, err := http.DefaultClient.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		health.Down().AddInfo("code", http.StatusBadRequest)

		return health
	}

	if resp.StatusCode != http.StatusOK {
		health.Down()
	}

	health.AddInfo("code", resp.StatusCode)

	return health
}
