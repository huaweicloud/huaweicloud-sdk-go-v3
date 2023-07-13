package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyRequest Request Object
type ListPolicyRequest struct {

	// instance id
	InstanceId string `json:"instance_id"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// marker
	Marker *string `json:"marker,omitempty"`

	// 是否查询上一页
	ReversePage *bool `json:"reverse_page,omitempty"`
}

func (o ListPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyRequest struct{}"
	}

	return strings.Join([]string{"ListPolicyRequest", string(data)}, " ")
}
