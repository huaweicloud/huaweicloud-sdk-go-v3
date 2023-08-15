package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEdgeGroupNodeBindingRequest Request Object
type UpdateEdgeGroupNodeBindingRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 边缘节点组ID
	GroupId string `json:"group_id"`

	Body *UpdateGroupNodeBindingRequest `json:"body,omitempty"`
}

func (o UpdateEdgeGroupNodeBindingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeGroupNodeBindingRequest struct{}"
	}

	return strings.Join([]string{"UpdateEdgeGroupNodeBindingRequest", string(data)}, " ")
}
