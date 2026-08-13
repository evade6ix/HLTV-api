package webclient

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func StartWebclient() *rod.Browser {
	// Explicitly launch Chromium without the sandbox so the scraper works in
	// containers such as GitHub Codespaces, Railway and Docker.
	controlURL := launcher.New().
		Headless(true).
		NoSandbox(true).
		MustLaunch()

	return rod.New().
		ControlURL(controlURL).
		Timeout(time.Minute).
		MustConnect()
}
