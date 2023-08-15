package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchImportSmnInfoReq 录入、修改收件方式与信息请求体。
type BatchImportSmnInfoReq struct {

	// 任务信息
	Jobs []SelectedSetAlarmTaskReq `json:"jobs"`

	AlarmNotifyInfo *BatchSetAlarmNotifyInfo `json:"alarm_notify_info"`
}

func (o BatchImportSmnInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchImportSmnInfoReq struct{}"
	}

	return strings.Join([]string{"BatchImportSmnInfoReq", string(data)}, " ")
}
