package model

type IpData struct {
	Ip       string `json:"ip"`
	HostName string `json:"hostname"`
	Active   bool   `json:"active"`
}

type IpConfig struct {

	Data     []IpData
}