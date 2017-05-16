package igsearch

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"
)

// UserSearch  ....
func UserSearch(query string) (IGSearchResult, error) {
	//url := "https://www.instagram.com/web/search/topsearch/?context=user&query=" + string + "&rank_token=0.6885969395144884"
	query = strings.Trim(query, "　 .\n\r\t&^%$$#@!*()_+-=/,[]{}<>?")
	result := IGSearchResult{}
	q := url.QueryEscape(query)
	if q == "" {
		return result, fmt.Errorf("need search query")
	}
	urlstr := fmt.Sprintf("https://www.instagram.com/web/search/topsearch/?context=user&query=%s", url.QueryEscape(query))
	ua := "GoogleBot"
	body, err := GetHTML(urlstr, ua)

	//log.Print(body)
	if err != nil {
		log.Print(err)
		return result, err
	}
	//parse JSON
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		log.Print(err)
		log.Print(body)
		return result, err
	}
	if result.Status != "ok" {
		log.Print("status is not ok")
		return result, fmt.Errorf("status is not ok")
	}
	return result, nil

}
