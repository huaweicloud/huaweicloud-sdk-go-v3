package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchAntivirusTaskRequestInfo 取消扫描任务
type SwitchAntivirusTaskRequestInfo struct {

	// 任务ID
	TaskId string `json:"task_id"`

	// 任务名称
	TaskName string `json:"task_name"`

	// 关联主机列表
	HostIds []string `json:"host_ids"`
}

func (o SwitchAntivirusTaskRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchAntivirusTaskRequestInfo struct{}"
	}

	return strings.Join([]string{"SwitchAntivirusTaskRequestInfo", string(data)}, " ")
}
