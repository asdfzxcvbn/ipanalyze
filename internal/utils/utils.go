package utils

import (
	"archive/zip"
	"errors"
	"strings"
)

// finds the name of the Info.plist in an ipa file.
func FindPlist(files []*zip.File) (*zip.File, error) {
	for _, file := range files {
		if strings.HasSuffix(file.Name, ".app/Info.plist") {
			return file, nil
		}
	}

	return nil, errors.New("couldn't find plist in files")
}
