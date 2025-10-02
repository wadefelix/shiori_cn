package core

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/go-shiori/shiori/internal/database"
)

var userAgent = "Shiori/2.0.0 (+https://github.com/go-shiori/shiori)"
var siteCookies = make(map[string]string)

func get_site_cookie(site_host string) (string,bool) {
	if cookie,ok := siteCookies[site_host]; ok {
		return cookie, ok
	}
	return "", false
}

func ReadSiteCookiesFromDB(db database.DB) error {
	if content,err := db.GetSiteCookies(context.TODO()); err==nil {
		var payload map[string]string
        err_j := json.Unmarshal([]byte(content), &payload)
        if err_j == nil {
			for k, v := range payload {
				if strings.Contains(k,",") {
					for _,kk := range strings.Split(k,",") {
				    siteCookies[kk] = v
					}
				} else {
				    siteCookies[k] = v
				}
			}
        } else {
			fmt.Errorf("Error during Unmarshal site-cookies: %v",  err_j)
		}
	}
	return nil
}
func init() {
    ua := os.Getenv("USER_AGENT")
    if len(ua) > 0 {
	    userAgent = ua
    }
}