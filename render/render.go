package render

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func check(err error) {
	if err != nil {
		fmt.Printf("Error : %s\n", err.Error())
		os.Exit(1)
	}
}

func copyFile(src string, dest string) {
	srcFile, err := os.Open(src)
	check(err)
	defer srcFile.Close()

	destFile, err := os.Create(dest) // creates if file doesn't exist
	check(err)
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
	check(err)

	err = destFile.Sync()
	check(err)
}

func Render(replayPath string) string {

	copyFile(replayPath, "./replays/replay.osr")

	fmt.Println("copied file")

	existingFiles, err := os.ReadDir("./danser/videos")
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("./danser/danser-cli", "-record", "-replay=./replays/replay.osr")
	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}

	stderr, _ := cmd.StderrPipe()
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		m := scanner.Text()
		if strings.Contains(m, "ETA") {
			fmt.Println(m)
		}
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Replay Done")

	err = os.Remove("./replays/replay.osr")
	if err != nil {
		log.Fatal(err)
	}

	newFiles, err := os.ReadDir("./danser/videos")
	if err != nil {
		log.Fatal(err)
	}

	outputFile := ""

	for _, f1 := range newFiles {
		found := false
		for _, f2 := range existingFiles {
			if f1 == f2 {
				found = true
			}
		}

		if !found {
			outputFile = f1.Name()
		}
	}

	return outputFile

}
