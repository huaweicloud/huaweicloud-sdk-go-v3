package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchNodeConnectionRequest Request Object
type SwitchNodeConnectionRequest struct {

	// 设备ID，从专业版HiLens控制台设备管理[查询设备列表](ListNodeUsingGeT.xml)获取
	NodeId string `json:"node_id"`

	// 服务提供者：ief或hilens，选择设备纳管到不同的平台。不填默认为hilens平台
	Provider *string `json:"provider,omitempty"`

	// 设备启用/停用动作，启用（start）,停用（stop）
	Action string `json:"action"`
}

func (o SwitchNodeConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchNodeConnectionRequest struct{}"
	}

	return strings.Join([]string{"SwitchNodeConnectionRequest", string(data)}, " ")
}
