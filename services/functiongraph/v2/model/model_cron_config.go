package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CronConfig 定时配置
type CronConfig struct {

	// 定时配置名称
	Name *string `json:"name,omitempty"`

	// 定时表达式
	Cron *string `json:"cron,omitempty"`

	// 拉起预留实例个数
	Count *int32 `json:"count,omitempty"`

	// 开始时间戳
	StartTime *int64 `json:"start_time,omitempty"`

	// 失效时间戳
	ExpiredTime *int64 `json:"expired_time,omitempty"`
}

func (o CronConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CronConfig struct{}"
	}

	return strings.Join([]string{"CronConfig", string(data)}, " ")
}
