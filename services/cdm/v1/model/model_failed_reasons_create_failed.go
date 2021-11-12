package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群创建失败原因。
type FailedReasonsCreateFailed struct {
	// 错误码

	ErrorCode *string `json:"errorCode,omitempty"`
	// 失败原因

	ErrorMsg *string `json:"errorMsg,omitempty"`
}

func (o FailedReasonsCreateFailed) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailedReasonsCreateFailed struct{}"
	}

	return strings.Join([]string{"FailedReasonsCreateFailed", string(data)}, " ")
}
