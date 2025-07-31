package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityGroupPoliciesResponse Response Object
type ListSecurityGroupPoliciesResponse struct {

	// 总共记录数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 数据列表
	DataList *[]SecurityGroupPolicyResponseInfo `json:"data_list,omitempty"`

	// 数据最近同步时间
	LastUpdateTime *int64 `json:"last_update_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSecurityGroupPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityGroupPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityGroupPoliciesResponse", string(data)}, " ")
}
