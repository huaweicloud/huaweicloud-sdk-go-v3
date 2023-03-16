package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePolicy 操作的请求体。
type UpdatePolicyReqBody struct {

	// 要添加到新策略的策略文本内容。
	Content *string `json:"content,omitempty"`

	// 要分配给策略的可选说明。
	Description *string `json:"description,omitempty"`

	// 要分配给策略的名称。
	Name *string `json:"name,omitempty"`
}

func (o UpdatePolicyReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyReqBody struct{}"
	}

	return strings.Join([]string{"UpdatePolicyReqBody", string(data)}, " ")
}
