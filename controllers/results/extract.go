package results

import (
	"fmt"
	"hltv/models"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ExtractInfFromHTML(html string) (data []models.Results, err error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	var matches []models.Results
	matches = []models.Results{}
	doc.Find(".results-sublist").Each(func(i int, s *goquery.Selection) {
		headline := strings.Replace(s.Find(".standard-headline").Text(), "Results for", "", 1)
		headline = strings.TrimSpace(headline)
		fmt.Println("Group:", headline)
		linkTag := s.Find("a.a-reset").First()
		matchUrl, exists := linkTag.Attr("href")
		if !exists {
			matchUrl = ""
		}
		team1Name := s.Find(".team1 .team").First().Text()
		team2Name := s.Find(".team2 .team").First().Text()

		wonScoreStr := s.Find(".score-won").First().Text()
		lostScoreStr := s.Find(".score-lost").First().Text()

		wonScore, _ := strconv.Atoi(wonScoreStr)
		lostScore, _ := strconv.Atoi(lostScoreStr)

		event := s.Find(".event-name").First().Text()
		bo := s.Find(".map-text").First().Text()

		match := models.Results{
			MatchUrl: fmt.Sprintf("https://www.hltv.org/%v", matchUrl),
			Team1: models.Team{
				Name:    team1Name,
				WonMaps: wonScore,
			},
			Date: headline,
			Team2: models.Team{
				Name:    team2Name,
				WonMaps: lostScore,
			},
			Event: event,
			BO:    bo,
		}

		matches = append(matches, match)
	})

	return matches, nil
}
