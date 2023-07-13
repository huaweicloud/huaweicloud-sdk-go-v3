package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEdgeGroupRequest Request Object
type UpdateEdgeGroupRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 边缘节点组ID
	GroupId string `json:"group_id"`

	Body *EdgeGroupUpdateRequest `json:"body,omitempty"`
}

func (o UpdateEdgeGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateEdgeGroupRequest", string(data)}, " ")
}
