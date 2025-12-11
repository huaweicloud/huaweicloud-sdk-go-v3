package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchAntivirusTaskRequestInfo 取消扫描任务
type SwitchAntivirusTaskRequestInfo struct {

	// **参数解释**： 任务ID **取值范围**: 字符长度1-64位
	TaskId string `json:"task_id"`

	// **参数解释**: 任务名称 **取值范围**: 最大长度255个unicode字符。
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
