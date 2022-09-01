package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MasterEipResponseStatus struct {

	// 集群访问的PrivateIP(HA集群返回VIP)
	PrivateEndpoint *string `json:"privateEndpoint,omitempty" xml:"privateEndpoint"`

	// 集群访问的PublicIP
	PublicEndpoint *string `json:"publicEndpoint,omitempty" xml:"publicEndpoint"`
}

func (o MasterEipResponseStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MasterEipResponseStatus struct{}"
	}

	return strings.Join([]string{"MasterEipResponseStatus", string(data)}, " ")
}
