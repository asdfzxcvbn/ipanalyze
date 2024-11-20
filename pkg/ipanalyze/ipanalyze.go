package ipanalyze

import (
	"archive/zip"
	"io"
	"os"
	"path"

	"github.com/asdfzxcvbn/ipanalyze/internal/utils"
	"howett.net/plist"
)

// takes a path to a .ipa file and returns info.
func AnalyzeIPA(ipa string) (*IPAInfo, error) {
	o, err := zip.OpenReader(ipa)
	if err != nil {
		return nil, err
	}
	defer o.Close()

	plistFile, err := utils.FindPlist(o.File)
	if err != nil {
		return nil, err
	}

	po, err := plistFile.Open()
	if err != nil {
		return nil, err
	}
	defer po.Close()

	plistContents, err := io.ReadAll(po)
	if err != nil {
		return nil, err
	}

	var info IPAInfo
	if _, err = plist.Unmarshal(plistContents, &info); err != nil {
		return nil, err
	}

	return &info, nil
}

// takes a path to a .app folder and returns info.
func AnalyzeApp(app string) (*IPAInfo, error) {
	plistPath := path.Join(app, "Info.plist")

	file, err := os.Open(plistPath)
	if err != nil {
		return nil, err
	}
	defer file.Close() // oh well, wont check for error

	contents, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var info IPAInfo
	if _, err := plist.Unmarshal(contents, &info); err != nil {
		return nil, err
	}

	return &info, nil
}
