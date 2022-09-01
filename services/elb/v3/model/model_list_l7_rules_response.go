package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListL7RulesResponse struct {

	// 请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty" xml:"page_info"`

	// 规则对象列表。
	Rules          *[]L7Rule `json:"rules,omitempty" xml:"rules"`
	HttpStatusCode int       `json:"-"`
}

func (o ListL7RulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListL7RulesResponse struct{}"
	}

	return strings.Join([]string{"ListL7RulesResponse", string(data)}, " ")
}
