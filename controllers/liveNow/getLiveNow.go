package livenow

import (
	"fmt"
	"hltv/models"
	webclient "hltv/webClient"

	"github.com/go-rod/stealth"
)

func ExtractLiveNow(date string) (list []models.Match, err error) {
	web := webclient.StartWebclient()
	defer web.MustClose()
	page := stealth.MustPage(web)
	page.MustNavigate(fmt.Sprintf("https://www.hltv.org/matches?selectedDate=%v", date))
	el := page.MustElement("body > div.bgPadding > div.widthControl > div:nth-child(2) > div.contentCol > div.matches-v4 > div.new-standardPageGrid > div.mainContent > div.matches-list-column > div.matches-chronologically > div > div")
	html := string(el.MustHTML())
	data, err := ExtractInfFromHTML(html)
	if err != nil {
		return nil, nil
	}
	return data, nil

}
