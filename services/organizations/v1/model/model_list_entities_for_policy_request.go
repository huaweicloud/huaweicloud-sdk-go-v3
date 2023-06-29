package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEntitiesForPolicyRequest Request Object
type ListEntitiesForPolicyRequest struct {

	// 策略的唯一标识符（ID）。
	PolicyId string `json:"policy_id"`

	// 页面中最大结果数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListEntitiesForPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEntitiesForPolicyRequest struct{}"
	}

	return strings.Join([]string{"ListEntitiesForPolicyRequest", string(data)}, " ")
}
