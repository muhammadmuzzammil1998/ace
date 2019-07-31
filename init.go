package main

import (
	"flag"
	"os"
	"runtime"
)

func init() {
	flag.Float64Var(&flags.interval, "interval", 5, "Specifies the interval in which ace should crawl in minutes")

	flag.StringVar(&flags.datafile, "r", ".ace-resources", "Resource data file\nCan be paired with: {-add, -interval, -remove}")
	flag.StringVar(&flags.url, "u", "", "URL\nCan be paired with: {-add, -r, -remove}")
	flag.StringVar(&flags.action, "action", "", "What to trigger?\nOptions: {webhook, script}\nCan be paired with: {-add, -f, -location, -u}")
	flag.StringVar(&flags.location, "location", "", "Location of trigger set by action\nCan be paired with: {-action, -add, -f, -u}\nExamples:\n webhook: https://example.com/hook, https://example.com/hook.php\n script: /bin/exec, ~/ace-script.sh")

	flag.BoolVar(&flags.add, "add", false, "Add a resource to crawl\nCan be paired with: {-f, -u, -action, -location}")
	flag.BoolVar(&flags.version, "version", false, "Print ace version")
	flag.BoolVar(&flags.forceAdd, "f", false, "Skip test and force add a resource to crawl\nOnly to be paired with: {-add}")
	flag.BoolVar(&flags.remove, "remove", false, "Remove a resource\nCan be paired with: {-r, -u}")

	flag.Parse()

	user, err := os.UserHomeDir()
	check(err)
	sep := "/"
	if runtime.GOOS == "windows" {
		sep = "\\"
	}
	home = user + sep + ".ace" + sep

	flags.datafile = home + flags.datafile

	version = "1.19.7.1" // Major.Year.Month.Release
}
