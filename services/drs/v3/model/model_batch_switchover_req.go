package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSwitchoverReq 批量主备倒换请求体
type BatchSwitchoverReq struct {

	// 批量主备调换的任务详情ID请求列表
	Jobs []string `json:"jobs"`

	// 灾备倒换后是否暂停任务。默认值为false，不暂停灾备任务。
	IsPauseJob *bool `json:"is_pause_job,omitempty"`
}

func (o BatchSwitchoverReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSwitchoverReq struct{}"
	}

	return strings.Join([]string{"BatchSwitchoverReq", string(data)}, " ")
}
