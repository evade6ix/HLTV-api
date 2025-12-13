package webclient

import (
	"time"

	"github.com/go-rod/rod"
)

func StartWebclient() *rod.Browser {
	browser := rod.New().Timeout(time.Minute).MustConnect()
	return browser
}
