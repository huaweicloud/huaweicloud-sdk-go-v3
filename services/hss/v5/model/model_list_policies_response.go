package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPoliciesResponse Response Object
type ListPoliciesResponse struct {

	// 策略信息列表
	DataList *[]AgentPolicyInfo `json:"data_list,omitempty"`

	// 策略总数
	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListPoliciesResponse", string(data)}, " ")
}
