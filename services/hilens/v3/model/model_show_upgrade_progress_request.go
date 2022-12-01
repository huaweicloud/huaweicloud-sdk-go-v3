package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowUpgradeProgressRequest struct {

	// 设备ID，从专业版HiLens控制台设备管理[查询设备列表](ListNodeUsingGeT.xml)获取
	NodeId string `json:"node_id"`
}

func (o ShowUpgradeProgressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUpgradeProgressRequest struct{}"
	}

	return strings.Join([]string{"ShowUpgradeProgressRequest", string(data)}, " ")
}
