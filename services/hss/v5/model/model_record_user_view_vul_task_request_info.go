package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecordUserViewVulTaskRequestInfo struct {

	// 任务类型 - vul_repair：修复任务 - vul_scan：扫描任务
	TaskType string `json:"task_type"`
}

func (o RecordUserViewVulTaskRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordUserViewVulTaskRequestInfo struct{}"
	}

	return strings.Join([]string{"RecordUserViewVulTaskRequestInfo", string(data)}, " ")
}
