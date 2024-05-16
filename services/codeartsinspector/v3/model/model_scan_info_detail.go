package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScanInfoDetail struct {

	// 扫描任务创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 弱密码检查
	EnableWeakPasswd *bool `json:"enable_weak_passwd,omitempty"`

	// 扫描任务结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 任务进度
	Progress *int32 `json:"progress,omitempty"`

	// 任务描述
	Reason *string `json:"reason,omitempty"`

	// 扫描任务状态:   * 0 运行中   * 1 已完成   * 2 手动取消   * 3 等待中   * 4 扫描失败   * 5 等待定时调度
	Status *int32 `json:"status,omitempty"`
}

func (o ScanInfoDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScanInfoDetail struct{}"
	}

	return strings.Join([]string{"ScanInfoDetail", string(data)}, " ")
}
