package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePolicyRequestBody struct {

	// 策略名称（策略名称只能由数字、字母和下划线组成，长度不能超过64为字符）
	Name string `json:"name"`
}

func (o CreatePolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePolicyRequestBody", string(data)}, " ")
}
