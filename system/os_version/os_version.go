package os_version

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/StackExchange/wmi"
	"golang.org/x/sys/windows/registry"
)

// Fields from Win32_OperatingSystem WMI class, not all fields are used
type Win32_OperatingSystem struct {
	Caption        string
	Version        string
	OSArchitecture string
	InstallDate    time.Time
}

// OSVersion represents the operating system version information
type OSVersion struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	Major        uint64 `json:"major"`
	Minor        uint64 `json:"minor"`
	Patch        uint64 `json:"patch"`
	Build        string `json:"build"`
	Platform     string `json:"platform"`
	PlatformLike string `json:"platform_like"`
	Codename     string `json:"codename"`
	Arch         string `json:"arch"`
	InstallDate  int64  `json:"install_date"`
	Revision     uint64 `json:"revision"`
}

// GenOSVersion retrieves the Windows operating system version information
func GenOSVersion() (*OSVersion, error) {
	var winOS []Win32_OperatingSystem
	query := "SELECT Caption, Version, OSArchitecture, InstallDate FROM Win32_OperatingSystem"
	if err := wmi.Query(query, &winOS); err != nil {
		return nil, fmt.Errorf("failed to query WMI: %v", err)
	}

	// Create OSVersion struct
	osVersion := &OSVersion{
		Platform:     "windows",
		PlatformLike: "windows",
		Name:         winOS[0].Caption,
		Codename:     winOS[0].Caption,
		Version:      winOS[0].Version,
		Arch:         winOS[0].OSArchitecture,
		InstallDate:  winOS[0].InstallDate.Unix(),
	}

	// Parse version components
	parts := strings.Split(winOS[0].Version, ".")
	if len(parts) >= 1 {
		osVersion.Major, _ = strconv.ParseUint(parts[0], 10, 64)
	}
	if len(parts) >= 2 {
		osVersion.Minor, _ = strconv.ParseUint(parts[1], 10, 64)
	}
	if len(parts) >= 3 {
		osVersion.Build = parts[2]
	}

	// Get Windows-specific information from Registry
	k, err := registry.OpenKey(
		registry.LOCAL_MACHINE,
		`SOFTWARE\Microsoft\Windows NT\CurrentVersion`,
		registry.QUERY_VALUE)
	if err == nil {
		defer k.Close()

		// Get UBR (Update Build Revision)
		if ubr, _, err := k.GetIntegerValue("UBR"); err == nil {
			osVersion.Revision = ubr
		}

		// Get DisplayVersion for patch if available (Windows 10 and later)
		if displayVersion, _, err := k.GetStringValue("DisplayVersion"); err == nil {
			osVersion.Patch, _ = strconv.ParseUint(displayVersion, 10, 64)
		}
	}

	return osVersion, nil
}
