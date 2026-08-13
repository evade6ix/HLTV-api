package matches

import (
	"fmt"
	"hltv/models"
	webclient "hltv/webClient"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/stealth"
)

func ExtractMatches(date string) (list []models.Match, err error) {
	web := webclient.StartWebclient()
	defer web.MustClose()
	page := stealth.MustPage(web)
	url := fmt.Sprintf("https://www.hltv.org/matches?selectedDate=%v-%v-%v", time.Now().Year(), int(time.Now().Month()), date)
	err = rod.Try(func() {
		page.Timeout(30 * time.Second).MustNavigate(url)
		page.MustWaitLoad()
		page.MustElement("body")
	})
	if err != nil {
		return nil, fmt.Errorf("HLTV page did not load: %w", err)
	}
	html := page.MustHTML()
	data, err := ExtractInfFromHTML(html)
	if err != nil {
		return nil, nil
	}
	return data, nil
}
