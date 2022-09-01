package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量执行错误信息
type OperateErrorInfo struct {

	// 唯一标识ID，流程URN
	WorkflowUrn *string `json:"workflow_urn,omitempty" xml:"workflow_urn"`

	// 错误详情
	ErrorDetail *string `json:"error_detail,omitempty" xml:"error_detail"`
}

func (o OperateErrorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateErrorInfo struct{}"
	}

	return strings.Join([]string{"OperateErrorInfo", string(data)}, " ")
}
