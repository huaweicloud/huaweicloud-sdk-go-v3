package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNodeRequest Request Object
type UpdateNodeRequest struct {

	// 设备ID，从专业版HiLens控制台设备管理[查询设备列表](ListNodeUsingGeT.xml)获取
	NodeId string `json:"node_id"`

	Body *UpdateNodeRequestBody `json:"body,omitempty"`
}

func (o UpdateNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNodeRequest struct{}"
	}

	return strings.Join([]string{"UpdateNodeRequest", string(data)}, " ")
}
