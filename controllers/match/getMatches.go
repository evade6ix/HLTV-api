package match

import (
	"fmt"
	"hltv/models"
	webclient "hltv/webClient"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/stealth"
)

func ExtractMaps() (list []models.MapVeto, err error) {
	web := webclient.StartWebclient()
	defer web.MustClose()
	page := stealth.MustPage(web)
	url := fmt.Sprintf("https://www.hltv.org/matches/2388129/faze-vs-natus-vincere-starladder-budapest-major-2025")
	page.MustNavigate(url)
	err = rod.Try(func() {
		page.Timeout(5 * time.Second).MustElement("body")
	})
	if err != nil {
		return nil, fmt.Errorf("Body not load")
	}
	html := page.MustElement("body > div.bgPadding > div.widthControl > div:nth-child(2) > div.contentCol > div.match-page > div.g-grid.maps > div.col-6.col-7-small > div:nth-child(3) > div").MustHTML()
	data, err := ExtractMapsVet(html)
	if err != nil {
		return nil, nil
	}
	return data, nil
}
