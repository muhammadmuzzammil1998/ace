package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	if flags.version {
		fmt.Println("ace v" + version)
		fmt.Println("  by Muhammad Muzzammil.")
		fmt.Println("  https://muzzammil.xyz")
		return
	}
	checkdatafile()
	if flags.add {
		add()
		return
	}
	if flags.remove {
		remove()
		return
	}
	repeat(run, fmt.Sprintf("%f%s", flags.interval, "m"))
}

func ask(s string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(s + " ")
	r, _, _ := reader.ReadLine()
	return string(r)
}

func repeat(f func(), i string) {
	f()
	d, err := time.ParseDuration(i)
	check(err)
	for range time.Tick(d) {
		f()
	}
}

func uniqueSlice(slice ...string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func appendToken(str string, length int) string {
	set := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = set[seed.Intn(len(set))]
	}

	u, err := url.Parse(str)
	check(err)

	param := u.Query()
	param.Add("ace-token", string(randomString))

	u.RawQuery = param.Encode()
	return u.String()
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func verifyInput(s *string, h string) {
	if strings.Trim(*s, " ") == "" {
		*s = ask(h)
	}
}

func getTime() string {
	return time.Now().Format(time.RFC850)
}

func checkdatafile() bool {
	check(os.MkdirAll(home, os.ModeDir))
	if _, err := os.Stat(flags.datafile); err != nil {
		fmt.Printf("No data file found (%s).\n", flags.datafile)
		fmt.Println("Creating data file...")
		(&resources{}).write()
		return true
	}
	return false
}
