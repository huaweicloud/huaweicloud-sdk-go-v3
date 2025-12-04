package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDdmNodeDetailRequest Request Object
type ShowDdmNodeDetailRequest struct {

	// 实例 ID。
	InstanceId string `json:"instance_id"`

	// 节点 ID。
	NodeId string `json:"node_id"`
}

func (o ShowDdmNodeDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDdmNodeDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowDdmNodeDetailRequest", string(data)}, " ")
}
