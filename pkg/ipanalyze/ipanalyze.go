package ipanalyze

import (
	"archive/zip"
	"io"
	"os"
	"path"

	"github.com/asdfzxcvbn/ipanalyze/internal/ipas"
	"howett.net/plist"
)

// AnalyzeIPA takes a path to a .ipa file and returns info.
func AnalyzeIPA(ipa string) (*IPAInfo, error) {
	o, err := zip.OpenReader(ipa)
	if err != nil {
		return nil, err
	}
	defer o.Close()

	plistFile, err := ipas.FindPlist(o.File)
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

	// can't use decoder, no Seek method
	var info IPAInfo
	_, err = plist.Unmarshal(plistContents, &info)
	return &info, err
}

// AnalyzeApp takes a path to a .app directory and returns info.
func AnalyzeApp(app string) (*IPAInfo, error) {
	plistPath := path.Join(app, "Info.plist")

	file, err := os.Open(plistPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var info IPAInfo
	err = plist.NewDecoder(file).Decode(&info)
	return &info, err
}
