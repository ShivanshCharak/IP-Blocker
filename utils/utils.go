package utils

import "fmt"

func AddWWW(sites []string) (updatedSites []string) {
	updatedSites = sites
	for _, sites := range sites {
		prefixedSites := fmt.Sprintf("www.%s", sites)
		updatedSites = append(updatedSites, prefixedSites)
	}
	return
}
