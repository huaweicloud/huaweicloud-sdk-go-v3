package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSecurityGroupsRequest struct {
	// 查询返回边缘安全组列表数量。取值范围：0~1000。

	Limit *int32 `json:"limit,omitempty"`
	// 查询的偏移量。

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListSecurityGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityGroupsRequest", string(data)}, " ")
}
