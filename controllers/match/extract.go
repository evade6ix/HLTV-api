package match

import (
	"hltv/models"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ExtractMapsResultData(html string) (data []models.MapResult, err error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	doc.Find(".flexbox-column .mapholder").Each(func(i int, mapholder *goquery.Selection) {

		var result models.MapResult

		result.MapName = strings.TrimSpace(
			mapholder.Find(".mapname").First().Text(),
		)

		left := mapholder.Find(".results-left").First()
		result.Team1.TeamName = strings.TrimSpace(
			left.Find(".results-teamname").First().Text(),
		)

		if left.HasClass("won") {
			result.Team1.WonMaps = 1
		} else {
			result.Team1.WonMaps = 0
		}

		right := mapholder.Find(".results-right").First()
		result.Team2.TeamName = strings.TrimSpace(
			right.Find(".results-teamname").First().Text(),
		)

		if right.HasClass("won") {
			result.Team2.WonMaps = 1
		} else {
			result.Team2.WonMaps = 0
		}

		data = append(data, result)
	})

	return data, nil
}

func ExtractMapsVet(html string) (data []models.MapVeto, err error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	var mapChoices []models.MapVeto
	doc.Find(".padding").Each(func(i int, s *goquery.Selection) {
		s.Children().Each(func(j int, item *goquery.Selection) {
			text := strings.TrimSpace(item.Text())
			text = strings.ReplaceAll(text, "\n", " ")
			text = strings.Join(strings.Fields(text), " ")

			re := regexp.MustCompile(`^(\d+)\.\s+(.+?)\s+(removed|picked)\s+(.*)$`)
			match := re.FindStringSubmatch(text)

			if len(match) == 5 {
				mapChoices = append(mapChoices, models.MapVeto{
					TeamName:   match[2],
					TeamChoice: match[3],
					MapName:    match[4],
				})
				return
			}

			reLeft := regexp.MustCompile(`^(.*?)\s+was left over$`)
			left := reLeft.FindStringSubmatch(text)
			if len(left) == 2 {
				mapChoices = append(mapChoices, models.MapVeto{
					TeamChoice: "decider",
					MapName:    left[1],
				})
			}
		})
	})

	return mapChoices, nil
}
