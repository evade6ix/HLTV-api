package livenow

import (
	"fmt"
	"hltv/models"
	webclient "hltv/webClient"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/stealth"
)

func ExtractLiveNow(date string) (list []models.Match, err error) {
	web := webclient.StartWebclient()
	defer web.MustClose()
	page := stealth.MustPage(web)
	err = rod.Try(func() {
		page.Timeout(30 * time.Second).MustNavigate(fmt.Sprintf("https://www.hltv.org/matches?selectedDate=%v", date))
		page.MustWaitLoad()
		page.MustElement("body")
	})
	if err != nil {
		return nil, fmt.Errorf("HLTV page did not load: %w", err)
	}
	var emptyText string

	_ = rod.Try(func() {
		emptyText = page.
			MustElement("#for-you-empty-div").
			MustElement("b").
			MustText()
	})

	if emptyText != "" {
		return nil, fmt.Errorf("No matches running now")
	}
	var html string
	err = rod.Try(func() {
		html = page.Timeout(10 * time.Second).MustHTML()
	})
	if err != nil {
		// Live matches are optional. HLTV sometimes leaves the live page in a
		// loading state when no match is active; return an empty feed instead of
		// allowing a slow live request to fail the whole dashboard.
		return []models.Match{}, nil
	}
	data, err := ExtractInfFromHTML(html)
	if err != nil {
		return nil, nil
	}
	return data, nil

}
