package main

import (
	"log"
	"os"

	"github.com/akamensky/argparse"
	"github.com/awgh/metaloot/metaloot"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	parser := argparse.NewParser("metaloot", "another word for pirate treasure")
	basedir := parser.String("d", "dir", &argparse.Options{Required: false,
		Default: dir, Help: "Output Directory"})
	if err := parser.Parse(os.Args); err != nil {
		log.Println(parser.Usage(err))
		return
	}
	metaloot.Metaloot(*basedir, "http://169.254.169.254")
}
