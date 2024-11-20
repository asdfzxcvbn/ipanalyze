package ipanalyze

import "fmt"

type IPAInfo struct {
	BundleIdentifier string `plist:"CFBundleIdentifier"`
	MinimumOSVersion string `plist:"MinimumOSVersion"`

	Name        string `plist:"CFBundleName"`
	DisplayName string `plist:"CFBundleDisplayName"`

	Version      string `plist:"CFBundleVersion"`
	ShortVersion string `plist:"CFBundleShortVersionString"`
}

// print bundle id, name, version, and minimumosversion to stdout.
func (p IPAInfo) Print() {
	fmt.Println("[*] bundle id:", p.BundleIdentifier)

	if p.DisplayName != "" {
		fmt.Println("[*] display name:", p.DisplayName)
	} else if p.Name != "" {
		fmt.Println("[*] name:", p.Name)
	} else {
		fmt.Println("[?] no name found")
	}

	if p.ShortVersion != "" {
		fmt.Println("[*] short version:", p.ShortVersion)
	} else if p.Version != "" {
		fmt.Println("[*] version:", p.Version)
	} else {
		fmt.Println("[?] no version found")
	}

	fmt.Println("[*] minimum version:", p.MinimumOSVersion)
}
