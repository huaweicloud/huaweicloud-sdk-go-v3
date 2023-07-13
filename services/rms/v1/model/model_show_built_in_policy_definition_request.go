package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBuiltInPolicyDefinitionRequest Request Object
type ShowBuiltInPolicyDefinitionRequest struct {

	// 策略ID
	PolicyDefinitionId string `json:"policy_definition_id"`

	// 选择接口返回的信息的语言，默认为\"zh-cn\"中文
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowBuiltInPolicyDefinitionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBuiltInPolicyDefinitionRequest struct{}"
	}

	return strings.Join([]string{"ShowBuiltInPolicyDefinitionRequest", string(data)}, " ")
}
