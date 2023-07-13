package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PrecheckFailSubJobResult 失败子任务详情。
type PrecheckFailSubJobResult struct {

	// 子任务ID。
	Id *string `json:"id,omitempty"`

	// 子任务名称。
	Name *string `json:"name,omitempty"`

	// 子任务检查结果。
	CheckResult *string `json:"check_result,omitempty"`
}

func (o PrecheckFailSubJobResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrecheckFailSubJobResult struct{}"
	}

	return strings.Join([]string{"PrecheckFailSubJobResult", string(data)}, " ")
}
