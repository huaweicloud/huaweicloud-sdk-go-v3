package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoStop 训练作业的自动停止配置。
type AutoStop struct {

	// 时间单位。可选取值如下： - HOURS
	TimeUnit string `json:"time_unit"`

	// 运行时长，最小值为1。
	Duration int32 `json:"duration"`
}

func (o AutoStop) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoStop struct{}"
	}

	return strings.Join([]string{"AutoStop", string(data)}, " ")
}
