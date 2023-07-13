package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FreezeNodeRequest Request Object
type FreezeNodeRequest struct {

	// 设备ID，从专业版HiLens控制台设备管理[查询设备列表](ListNodeUsingGeT.xml)获取
	NodeId string `json:"node_id"`
}

func (o FreezeNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FreezeNodeRequest struct{}"
	}

	return strings.Join([]string{"FreezeNodeRequest", string(data)}, " ")
}
