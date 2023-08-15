package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerInfo 批量创建保护实例虚拟机信息
type ServerInfo struct {

	// 指定的生产站点云服务器ID。
	ServerId string `json:"server_id"`

	// 指定的容灾站点云服务器的flavor ID。
	FlavorRef *string `json:"flavorRef,omitempty"`
}

func (o ServerInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerInfo struct{}"
	}

	return strings.Join([]string{"ServerInfo", string(data)}, " ")
}
