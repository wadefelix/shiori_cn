package core
import (
	"os"
	"encoding/json"
    "io/ioutil"
	"fmt"
)

var userAgent = "Shiori/2.0.0 (+https://github.com/go-shiori/shiori)"
var siteCookies = make(map[string]string)

func init() {
    ua := os.Getenv("USER_AGENT")
    if len(ua) > 0 {
	    userAgent = ua
    }

	fileName := "/shiori/cookies.json"
	_, err := os.Stat(fileName)
	if err == nil {
		content, err := ioutil.ReadFile(fileName)
        if err == nil {
            var payload map[string]string
            err = json.Unmarshal(content, &payload)
            if err == nil {
				for k, v := range payload {
					siteCookies[k] = v
				}
            } else {
				fmt.Errorf("Error during Unmarshal(file %s): %v", fileName, err)
			}
        } else {
		    fmt.Errorf("cannot read %s", fileName)
		}
	} else {
		fmt.Errorf("%s not exists.", fileName)
	}
}