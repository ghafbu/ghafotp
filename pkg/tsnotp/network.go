package tsnotp

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
)

// تابع برای دریافت IP محلی
func getLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ip := ipnet.IP.String()
			if !strings.HasPrefix(ip, "169.254") { // نادیده گرفتن APIPA آدرس‌ها
				return ip, nil
			}
		}
	}
	return "", fmt.Errorf("no valid IP address found")
}

// تابع برای دریافت IP عمومی
func getPublicIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(ip), nil
}

// تابع برای دریافت MAC ID
func getMACID() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, i := range interfaces {
		if i.Flags&net.FlagUp != 0 && i.Flags&net.FlagLoopback == 0 {
			addrs, err := i.Addrs()
			if err != nil {
				return "", err
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
					return i.HardwareAddr.String(), nil
				}
			}
		}
	}
	return "", fmt.Errorf("no valid MAC address found")
}

// تابع واحد برای شمارش تعداد شبکه‌های در دسترس
// تابع واحد برای شمارش تعداد شبکه‌های در دسترس
func getNetworkCount() (int, error) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("netsh", "wlan", "show", "networks")
	case "linux":
		cmd = exec.Command("nmcli", "device", "wifi", "list")
	case "darwin":
		cmd = exec.Command("/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport", "--scan")
	case "android":
		cmd = exec.Command("sh", "-c", "dumpsys wifi | grep 'SSID'")
	case "ios":
		cmd = exec.Command("sh", "-c", "ifconfig en0 | grep 'inet '")
	default:
		return 0, fmt.Errorf("unsupported operating system")
	}

	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(output), "\n")
	count := len(lines) - 1

	return count, nil
}

// تابع برای دریافت اطلاعات شبکه
func getNetworkInfo() (string, int, string, error) {
	var networkCount int
	var publicIP, macID string
	var err error

	// دریافت IP عمومی
	publicIP, err = getPublicIP()
	if err != nil {
		return "", 0, "", err
	}

	// دریافت MAC ID
	macID, err = getMACID()
	if err != nil {
		return "", 0, "", err
	}

	// دریافت تعداد شبکه‌های در دسترس
	networkCount, err = getNetworkCount()
	if err != nil {
		networkCount = 0
	}

	return publicIP, networkCount, macID, nil
}
