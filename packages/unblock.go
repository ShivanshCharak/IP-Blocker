package block

import "github.com/txn2/txeh"

func CleanBlocks(hosts *txeh.Hosts, blacklistConfiguredSites []string) (cleanedTargets []string) {
	stateSites := blacklistConfiguredSites
	for _, stateSites := range stateSites {
		hosts.RemoveHost(stateSites)
		hosts.Save()
		cleanedTargets = append(cleanedTargets, stateSites)
	}
	return
}
