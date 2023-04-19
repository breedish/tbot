package service

import "fmt"

type Info struct {
	ServiceName string
	Port        string
	ApiVersion  string
	Env         string
}

func (info *Info) GetServiceURL() string {
	return fmt.Sprintf("/api/%s/%s", info.ServiceName, info.ApiVersion)
}
