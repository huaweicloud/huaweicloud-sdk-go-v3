package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityGroupRulesRequest Request Object
type ListSecurityGroupRulesRequest struct {

	// 查询返回边缘安全组规则列表数量。取值范围：0~1000。
	Limit *int32 `json:"limit,omitempty"`

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 安全组ID。uuid
	SecurityGroupId *string `json:"security_group_id,omitempty"`
}

func (o ListSecurityGroupRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityGroupRulesRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityGroupRulesRequest", string(data)}, " ")
}
