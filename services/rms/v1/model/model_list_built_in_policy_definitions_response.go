package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListBuiltInPolicyDefinitionsResponse struct {
	// 策略定义列表

	Value *[]PolicyDefinition `json:"value,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListBuiltInPolicyDefinitionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBuiltInPolicyDefinitionsResponse struct{}"
	}

	return strings.Join([]string{"ListBuiltInPolicyDefinitionsResponse", string(data)}, " ")
}
