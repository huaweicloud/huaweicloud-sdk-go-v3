package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNodeRequest Request Object
type DeleteNodeRequest struct {

	// 设备ID，从专业版HiLens控制台设备管理[查询设备列表](ListNodeUsingGeT.xml)获取
	NodeId string `json:"node_id"`

	// 是否强制删除，true代表是，false代表否
	ForceDelete *bool `json:"force_delete,omitempty"`
}

func (o DeleteNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNodeRequest struct{}"
	}

	return strings.Join([]string{"DeleteNodeRequest", string(data)}, " ")
}
