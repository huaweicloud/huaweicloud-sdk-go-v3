package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteEdgeGroupRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 边缘节点组ID
	GroupId string `json:"group_id"`
}

func (o DeleteEdgeGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEdgeGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteEdgeGroupRequest", string(data)}, " ")
}
