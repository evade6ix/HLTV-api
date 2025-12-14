package match

import (
	"hltv/models"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

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
