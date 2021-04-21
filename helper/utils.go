package helper

import (
	"strings"
)

var (
	tarVersion = ""
	// IsGnuTar show tar type
	IsGnuTar = false
	IsBusyBoxTar = false
)

func init() {
	getTarVersion()
	checkIsGnuTar()
	checkIsBusyBoxTar()
}

func getTarVersion() {
	out, _ := Exec("tar", "--version")
	tarVersion = out
}

func checkIsGnuTar() {
	IsGnuTar = strings.Contains(tarVersion, "GNU")
}

func checkIsBusyBoxTar() {
	IsBusyBoxTar = strings.Contains(tarVersion, "busybox")
}

// CleanHost clean host url ftp://foo.bar.com -> foo.bar.com
func CleanHost(host string) string {
	// ftp://ftp.your-host.com -> ftp.your-host.com
	if strings.Contains(host, "://") {
		return strings.Split(host, "://")[1]
	}

	return host
}
