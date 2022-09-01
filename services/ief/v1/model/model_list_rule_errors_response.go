package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRuleErrorsResponse struct {

	// 满足条件的错误个数
	Count *int64 `json:"count,omitempty" xml:"count"`

	// 错误列表
	Errors         *[]Error `json:"errors,omitempty" xml:"errors"`
	HttpStatusCode int      `json:"-"`
}

func (o ListRuleErrorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRuleErrorsResponse struct{}"
	}

	return strings.Join([]string{"ListRuleErrorsResponse", string(data)}, " ")
}
