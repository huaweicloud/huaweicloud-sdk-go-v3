package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstallStep 安装步骤详情。
type InstallStep struct {

	// 安装步骤。
	Name *string `json:"name,omitempty"`

	Status *StepStatus `json:"status,omitempty"`

	// 失败原因。
	FailedReason *string `json:"failed_reason,omitempty"`
}

func (o InstallStep) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallStep struct{}"
	}

	return strings.Join([]string{"InstallStep", string(data)}, " ")
}
