package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateHttpPolicyRequestBody struct {

	// 防护策略名
	Name string `json:"name"`
}

func (o CreateHttpPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHttpPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateHttpPolicyRequestBody", string(data)}, " ")
}
