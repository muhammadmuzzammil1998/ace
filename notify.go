package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func notify(d data, oldHash, newHash string) {
	fmt.Println("  Action:", d.Action)
	fmt.Println("   -- Trigger:", d.Location)

	if strings.Trim(d.Action, " ") == "" {
		fmt.Println("\n  No action found... skipping.")
		return
	}
	if strings.Trim(d.Location, " ") == "" {
		fmt.Println("\n  No trigger location found... skipping.")
		return
	}
	if _, err := os.Stat(flags.datafile); err != nil {
		fmt.Println("\n  Trouble with trigger location... skipping.")
		return
	}

	src, err := d.getSource()
	check(err)
	respdata := &response{
		Version: version,
		Data: responseData{
			NewHash:     newHash,
			OldHash:     oldHash,
			URL:         d.URL,
			Source:      src,
			LastUpdated: d.LastUpdated,
			IsForced:    d.IsForced,
		},
	}
	if d.Action == "webhook" {
		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		encoder.SetIndent(" ", " ")
		check(encoder.Encode(respdata))
		resp, err := http.Post(d.Location, "application/json", bytes.NewBuffer(buffer.Bytes()))
		check(err)
		fmt.Println("   -- Response:", resp.Status)
	} else if d.Action == "script" {
		var cmd *exec.Cmd
		temp, err := ioutil.TempFile(os.TempDir(), "ace-"+respdata.Data.NewHash)
		check(err)
		defer temp.Close()
		forced := "false"
		if respdata.Data.IsForced {
			forced = "true"
		}
		args := fmt.Sprintf("%s %s %s %s %s %s %s %s", d.Location, temp.Name(), respdata.Data.OldHash, respdata.Data.NewHash, respdata.Data.URL, respdata.Data.LastUpdated, forced, respdata.Version)
		if runtime.GOOS == "windows" {
			cmd = exec.Command("cmd.exe", "/c", args)
		} else {
			cmd = exec.Command("sh", "-c", args)
		}
		temp.WriteString(respdata.Data.Source)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("  ", err)
		}
		fmt.Printf("   -- Source saved at: %s\n", temp.Name())
		fmt.Printf("   ---Output start---\n%s\n", out)
		fmt.Printf("   ---Output end---\n")
	} else {
		check(errors.New("  Unknown action - " + d.Action))
	}
}
