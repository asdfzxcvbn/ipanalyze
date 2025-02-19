package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/asdfzxcvbn/ipanalyze/pkg/ipanalyze"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("usage: appanalyze <ipa/app>")
	}

	inputArg := os.Args[1]
	if _, err := os.Stat(inputArg); os.IsNotExist(err) {
		log.Fatalln(err)
	}
	inputBasename := filepath.Base(inputArg)

	var (
		info *ipanalyze.IPAInfo
		err  error
	)

	if strings.HasSuffix(inputBasename, ".ipa") {
		info, err = ipanalyze.AnalyzeIPA(inputArg)
	} else if strings.HasSuffix(inputBasename, ".app") {
		info, err = ipanalyze.AnalyzeApp(inputArg)
	} else {
		log.Fatalln("error: input must be either a .ipa file or a .app directory")
	}

	if err != nil {
		log.Fatalln(err)
	}

	info.Print()
}
