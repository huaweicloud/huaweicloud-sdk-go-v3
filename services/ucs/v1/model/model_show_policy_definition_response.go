package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPolicyDefinitionResponse Response Object
type ShowPolicyDefinitionResponse struct {

	// API类型，固定值“ConstraintTemplate”
	Kind *string `json:"kind,omitempty"`

	// API版本
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *ObjectMeta `json:"metadata,omitempty"`

	// 约束模板的详细属性
	Spec           *interface{} `json:"spec,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowPolicyDefinitionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPolicyDefinitionResponse struct{}"
	}

	return strings.Join([]string{"ShowPolicyDefinitionResponse", string(data)}, " ")
}
