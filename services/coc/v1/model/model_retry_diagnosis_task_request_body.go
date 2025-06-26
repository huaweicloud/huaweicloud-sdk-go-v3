package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RetryDiagnosisTaskRequestBody struct {

	// 被重试实例的ID
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o RetryDiagnosisTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryDiagnosisTaskRequestBody struct{}"
	}

	return strings.Join([]string{"RetryDiagnosisTaskRequestBody", string(data)}, " ")
}
