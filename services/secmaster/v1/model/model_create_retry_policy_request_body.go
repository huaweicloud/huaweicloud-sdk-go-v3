package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRetryPolicyRequestBody 创建请求策略请求body体
type CreateRetryPolicyRequestBody struct {
	DataObject *CreateRetryPolicyRequestBodyDataObject `json:"data_object,omitempty"`
}

func (o CreateRetryPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRetryPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRetryPolicyRequestBody", string(data)}, " ")
}
