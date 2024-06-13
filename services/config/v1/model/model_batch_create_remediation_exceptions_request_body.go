package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateRemediationExceptionsRequestBody 批量创建合规规则修正例外的详情。
type BatchCreateRemediationExceptionsRequestBody struct {

	// 批量创建合规规则修正例外的详情。
	Exceptions []RemediationExceptionRequest `json:"exceptions"`
}

func (o BatchCreateRemediationExceptionsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateRemediationExceptionsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateRemediationExceptionsRequestBody", string(data)}, " ")
}
