package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"runtime"
)

func convertAndClean(fname []byte) []byte {
	dashRegex := regexp.MustCompile(`_-_`)
	fname = dashRegex.ReplaceAll(fname, []byte(" - "))
	copiedRegex := regexp.MustCompile(`\([0-9]\)`)
	fname = copiedRegex.ReplaceAll(fname, nil)
	stupidWebsiteNamesRegex := regexp.MustCompile(`\(.*\..*\)`)
	fname = stupidWebsiteNamesRegex.ReplaceAll(fname, nil)
	mp3Regex := regexp.MustCompile(`_?\.mp3$`)
	fname = mp3Regex.ReplaceAll(fname, nil)

	return append(fname, 10)
}

func main() {
	fmt.Printf("NEWLINE CHARACTER: %d\n", int(byte("\n"[0])))
	mp3Regex := regexp.MustCompile(`.*\.mp3$`) //(`\.[a-zA-Z0-9]*?$`)
	outputFilename := "song_list"
	var slash string
	if runtime.GOOS != "windows" {
		slash = "\\"
	} else {
		slash = "/"
	}
	songs := []byte{}
	args := os.Args //arg1 is folder to read
	if len(args) < 2 {
		fmt.Print("ya goofd")
		return
	}
	files, err := ioutil.ReadDir(args[1])
	if err != nil {
		fmt.Println("Error: Unable to read directory")
	}
	for _, file := range files {
		fname := []byte(file.Name())
		if mp3Regex.Match(fname) {
			songs = append(songs, convertAndClean(fname)...)
		}
	}

	//WRITING
	outputFile, err := os.Create(args[1] + slash + outputFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes, err := outputFile.Write(songs)
	if err != nil {
		fmt.Println(err)
		outputFile.Close()
		return
	}
	fmt.Println(bytes, "bytes written successfully")
	//err =
	if outputFile.Close() != nil {
		fmt.Println(err)
		return
	}
}
