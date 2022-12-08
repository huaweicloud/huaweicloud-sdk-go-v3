package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAccessPoliciesRequest struct {

	// 每页数量,范围0-100,默认100。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量,范围0-99,默认0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListAccessPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessPoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListAccessPoliciesRequest", string(data)}, " ")
}
