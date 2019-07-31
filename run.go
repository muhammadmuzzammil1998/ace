package main

import (
	"errors"
	"fmt"

	pagehash "muzzammil.xyz/pagehashgo"
)

func run() {
	res := (&resources{}).read()
	if len(res.Data) < 1 {
		check(errors.New("Resource file is empty"))
		return
	}
	res.LastRun.Start = getTime()
	fmt.Printf("Started: %s\n", res.LastRun.Start)

	for i, data := range res.Data {
		fmt.Printf("Checking for %s", data.URL)
		oldHash := data.Hash
		hashes, err := pagehash.Get(appendToken(data.URL, 8))
		check(err)

		if newHash := hashes.GetSHA256(); oldHash != newHash {
			fmt.Printf(" - New hash found - %s\n", newHash)
			res.Data[i].Hash = newHash
			res.Data[i].LastUpdated = getTime()
			res.overwrite()
			notify(data, oldHash, newHash)
			continue
		}
		fmt.Println(" - Didn't change")
	}
	res.write()
	res.LastRun.End = getTime()
	fmt.Printf("Ended: %s\n\n--\n\n", res.LastRun.End)

	res.overwrite()
}
