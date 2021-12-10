package adventofcode2021

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

func DecodeInput(folder string) string {
	currentPath, _ := os.Getwd()

	file, err := os.Open(
		path.Join(
			currentPath,
			"/"+folder+"/",
			"/input.txt",
		),
	)

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	return string(b)
}
