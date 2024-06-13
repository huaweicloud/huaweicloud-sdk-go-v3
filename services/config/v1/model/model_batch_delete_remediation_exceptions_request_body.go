package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteRemediationExceptionsRequestBody 批量删除合规规则修正例外的详情。
type BatchDeleteRemediationExceptionsRequestBody struct {

	// 批量删除合规规则修正例外的详情。
	Exceptions []RemediationExceptionRequest `json:"exceptions"`
}

func (o BatchDeleteRemediationExceptionsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRemediationExceptionsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteRemediationExceptionsRequestBody", string(data)}, " ")
}
