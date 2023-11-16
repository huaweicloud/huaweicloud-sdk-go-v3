package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DomainRealServerInfo 域名源站信息
type DomainRealServerInfo struct {

	// 源站类型
	RealServerType *int32 `json:"real_server_type,omitempty"`

	// 源站
	RealServers *string `json:"real_servers,omitempty"`
}

func (o DomainRealServerInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainRealServerInfo struct{}"
	}

	return strings.Join([]string{"DomainRealServerInfo", string(data)}, " ")
}
