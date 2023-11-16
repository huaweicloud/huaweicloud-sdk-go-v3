package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePolicyRequestBody 创建策略的请求体
type CreatePolicyRequestBody struct {

	// 策略名
	Name string `json:"name"`

	// 防护包id
	PackageId string `json:"package_id"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o CreatePolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePolicyRequestBody", string(data)}, " ")
}
