package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnfreezeNodeRequest Request Object
type UnfreezeNodeRequest struct {

	// 设备ID，从专业版HiLens控制台设备管理[查询设备列表](ListNodeUsingGeT.xml)获取
	NodeId string `json:"node_id"`

	Body *ActivateNodeRequestBody `json:"body,omitempty"`
}

func (o UnfreezeNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnfreezeNodeRequest struct{}"
	}

	return strings.Join([]string{"UnfreezeNodeRequest", string(data)}, " ")
}
