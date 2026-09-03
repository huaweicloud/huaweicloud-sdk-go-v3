package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OperateNotice struct {

	// 发送告警渠道
	AlertChannel *string `json:"alert_channel,omitempty"`

	// 操作通知 0 关闭 1开启
	Enable *string `json:"enable,omitempty"`

	// 通知组列表
	Groups *[]AlertGroup `json:"groups,omitempty"`

	// 通知类型列表
	OperateTypes *[]string `json:"operateTypes,omitempty"`
}

func (o OperateNotice) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateNotice struct{}"
	}

	return strings.Join([]string{"OperateNotice", string(data)}, " ")
}
