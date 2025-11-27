package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyDefinitionsResponse Response Object
type ListPolicyDefinitionsResponse struct {

	// 约束模板列表
	Items *[]UcsConstraintTemplate `json:"items,omitempty"`

	// API类型
	Kind *string `json:"kind,omitempty"`

	// API版本
	ApiVersion     *string `json:"apiVersion,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPolicyDefinitionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyDefinitionsResponse struct{}"
	}

	return strings.Join([]string{"ListPolicyDefinitionsResponse", string(data)}, " ")
}
