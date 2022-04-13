package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowBaremetalServerVolumeInfoRequest struct {
	// 裸金属服务器ID。可以从裸金属服务器控制台查询，或者通过调用7.3.4-查询裸金属服务器列表（OpenStack原生）API获取。

	ServerId string `json:"server_id"`
}

func (o ShowBaremetalServerVolumeInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBaremetalServerVolumeInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowBaremetalServerVolumeInfoRequest", string(data)}, " ")
}
