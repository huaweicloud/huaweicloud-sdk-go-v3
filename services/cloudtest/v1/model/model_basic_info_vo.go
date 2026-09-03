package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BasicInfoVo struct {

	// 采集间隔以秒为单位
	CollectInterval *int32 `json:"collect_interval,omitempty"`

	// 子任务名称
	SubTaskName *string `json:"sub_task_name,omitempty"`
}

func (o BasicInfoVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasicInfoVo struct{}"
	}

	return strings.Join([]string{"BasicInfoVo", string(data)}, " ")
}
