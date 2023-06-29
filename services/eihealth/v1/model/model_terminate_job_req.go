package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TerminateJobReq 终止作业请求体
type TerminateJobReq struct {

	// 是否强制终止,默认为false
	Force bool `json:"force"`
}

func (o TerminateJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TerminateJobReq struct{}"
	}

	return strings.Join([]string{"TerminateJobReq", string(data)}, " ")
}
