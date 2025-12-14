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
	page.MustNavigate(url)
	err = rod.Try(func() {
		page.Timeout(5 * time.Second).MustElement("body")
	})
	if err != nil {
		return nil, fmt.Errorf("Body not load")
	}
	html := page.MustHTML()
	data, err := ExtractInfFromHTML(html)
	if err != nil {
		return nil, nil
	}
	return data, nil
}

/*
func ExtractMatches(date string) (list []models.Match, err error) {
	web := webclient.StartWebclient()
	defer web.MustClose()

	ctx, _ := web.Incognito()
	page := stealth.MustPage(ctx)

	page.MustNavigate("https://www.hltv.org/matches")
	page.MustWaitLoad()
	page.MustWaitIdle()

	page.MustElement(
		fmt.Sprintf(`div.calendar-day[data-calendar-day="%s"]`, date),
	).MustClick()

	page.MustWaitIdle()

	el := page.MustElement(
		"div.matches-list-wrapper div.matches-chronologically",
	)

	html := el.MustHTML()

	data, err := ExtractInfFromHTML(html)
	if err != nil {
		return nil, err
	}

	return data, nil
}

*/
