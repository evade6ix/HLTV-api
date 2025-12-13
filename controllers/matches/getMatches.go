package matches

import (
	"fmt"
	"hltv/models"
	webclient "hltv/webClient"

	"github.com/go-rod/stealth"
)

func ExtractMatches(date string) (list []models.Match, err error) {
	web := webclient.StartWebclient()
	defer web.MustClose()

	ctx, _ := web.Incognito()
	page := stealth.MustPage(ctx)

	page.MustNavigate("https://www.hltv.org/matches")
	page.MustWaitLoad()
	page.MustWaitIdle()

	linkButton := fmt.Sprintf("body > div.bgPadding > div.widthControl > div:nth-child(2) > div.contentCol > div.matches-v4 > div.new-standardPageGrid > div.matches-sidebar-wrapper > div > div.first-month.matches-events-list-hide-smartphone > div > div > div.calendar-dates > div:nth-child(%v)", date)
	page.MustElement(linkButton).MustClick()

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
