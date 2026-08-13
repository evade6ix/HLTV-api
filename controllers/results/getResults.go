package results

import (
	"fmt"
	"hltv/models"
	webclient "hltv/webClient"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/stealth"
)

func ExtractResults() (list []models.Results, err error) {
	web := webclient.StartWebclient()
	defer web.MustClose()
	page := stealth.MustPage(web)
	url := fmt.Sprintf("https://www.hltv.org/results")
	err = rod.Try(func() {
		page.Timeout(30 * time.Second).MustNavigate(url)
		page.MustWaitLoad()
		page.MustElement("body")
	})
	if err != nil {
		return nil, fmt.Errorf("HLTV page did not load: %w", err)
	}

	// Parse the complete page. The extractor already scopes itself to result
	// lists, and this avoids breakage when HLTV moves the surrounding layout.
	html := page.MustHTML()
	data, err := ExtractInfFromHTML(html)
	if err != nil {
		return nil, nil
	}
	return data, nil
}
