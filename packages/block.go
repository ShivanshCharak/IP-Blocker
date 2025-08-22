package block

import (
	"fmt"
	"net/http"

	"github.com/txn2/txeh"
)

func BlockSites(hosts *txeh.Hosts, sites []string) (blacklistConfiguredSites []string) {
	for _, s := range sites {
		blacklistConfiguredSites = append(blacklistConfiguredSites, s)
	}
	for _, site := range sites {
		if isTargetAlive(site) {
			blockSite(hosts, site)
		}
	}
	return blacklistConfiguredSites
}
func isTargetAlive(url string) bool {
	resp, err := http.Get(fmt.Sprintf("https://%s", url))
	if err != nil {
		return true
	}
	if resp.StatusCode >= 200 && resp.StatusCode <= 500 {
		return true
	} else {
		fmt.Printf("Status code for %s is %d ignoring\n", url, resp.StatusCode)
		return false
	}
}

func blockSite(hosts *txeh.Hosts, site string) {
	target := "0.0.0.0"
	hosts.AddHost(target, site)
	hosts.Save()
}
