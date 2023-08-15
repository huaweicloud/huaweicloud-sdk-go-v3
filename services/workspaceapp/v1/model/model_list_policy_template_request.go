package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyTemplateRequest Request Object
type ListPolicyTemplateRequest struct {

	// 查询的偏移量
	Offset int32 `json:"offset"`

	// 查询的数量，值区间[1-100]
	Limit int32 `json:"limit"`

	// 根据策略模板名字过滤结果。
	PolicyGroupName *string `json:"policy_group_name,omitempty"`
}

func (o ListPolicyTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyTemplateRequest struct{}"
	}

	return strings.Join([]string{"ListPolicyTemplateRequest", string(data)}, " ")
}
