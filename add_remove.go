package main

import (
	"fmt"
	"log"

	pagehash "muzzammil.xyz/pagehashgo"
)

func add() {
	verifyInput(&flags.url, "Enter the URL to add:")

	if !flags.forceAdd {
		var hashes [5]string
		for i := 4; i >= 0; i-- {
			fmt.Printf("Testing URL: Generating hashes... %d\r", i+1)
			h, err := pagehash.Get(flags.url)
			check(err)
			hashes[i] = h.GetSHA256()
		}
		if len(uniqueSlice(hashes[0], hashes[1], hashes[2], hashes[3], hashes[4])) > 2 {
			fmt.Printf("Testing URL: Test failed - URL is not static\r\n")
			return
		}
		fmt.Printf("Testing URL: Test passed              \r\n")
	}

	verifyInput(&flags.action, "Action - [w]ebhook or [s]cript:")

	if flags.action[0] == byte('w') {
		flags.action = "webhook"
	} else if flags.action[0] == byte('s') {
		flags.action = "script"
	} else {
		log.Fatalln("Invalid option.")
		return
	}

	hashes, err := pagehash.Get(flags.url)
	check(err)

	verifyInput(&flags.location, "Location: ")

	(&resources{}).read().add(data{
		URL:         flags.url,
		Hash:        hashes.GetSHA256(),
		Action:      flags.action,
		Location:    flags.location,
		IsForced:    flags.forceAdd,
		LastUpdated: getTime(),
	}).write()

	fmt.Printf("%s added to %s", flags.url, flags.datafile)
}

func remove() {
	verifyInput(&flags.url, "Enter the URL to remove:")

	if r := (&resources{}).read(); r.remove(flags.url) {
		r.overwrite()
		fmt.Printf("%s removed from %s", flags.url, flags.datafile)
		return
	}
	fmt.Print("URL not found in data.")
}
