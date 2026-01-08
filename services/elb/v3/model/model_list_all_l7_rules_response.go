package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllL7RulesResponse Response Object
type ListAllL7RulesResponse struct {

	// **参数解释**：请求ID。  **取值范围**：由数字、小写字母和中划线（-）组成的字符串，自动生成。
	RequestId *string `json:"request_id,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// **参数解释**：规则对象列表。
	Rules          *[]L7Rule `json:"rules,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAllL7RulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllL7RulesResponse struct{}"
	}

	return strings.Join([]string{"ListAllL7RulesResponse", string(data)}, " ")
}
