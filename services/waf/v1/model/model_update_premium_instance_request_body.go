package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePremiumInstanceRequestBody 独享引擎操作
type UpdatePremiumInstanceRequestBody struct {

	// 独享引擎操作名称
	Action string `json:"action"`

	// 具体的请求体
	Params *[]string `json:"params,omitempty"`
}

func (o UpdatePremiumInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePremiumInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePremiumInstanceRequestBody", string(data)}, " ")
}
