package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CcrulesListInfoTagCondition 用户标识，当限速模式为other时，需要传该参数。根据Referer（自定义请求访问的来源）字段区分单个Web访问者
type CcrulesListInfoTagCondition struct {

	// 用户标识字段，其值固定为referer
	Category *string `json:"category,omitempty"`

	// 用户标识字段内容
	Contents *[]string `json:"contents,omitempty"`
}

func (o CcrulesListInfoTagCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CcrulesListInfoTagCondition struct{}"
	}

	return strings.Join([]string{"CcrulesListInfoTagCondition", string(data)}, " ")
}
