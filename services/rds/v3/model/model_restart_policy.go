package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestartPolicy struct {

	// 重启周期配置，在重启周期内的可维护时间窗进行重启，时区为UTC，取值范围： begin: 每月的第一天 middle：每月的15号 end：每月的最后一天 空列表，表示该功能关闭。
	Period *[]string `json:"period,omitempty"`

	// 所在时区相较于UTC时间的偏移量，取值范围：格式必须为+hh:mm，或者-hh:mm，例如+08:00。
	UtcOffset *string `json:"utc_offset,omitempty"`
}

func (o RestartPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartPolicy struct{}"
	}

	return strings.Join([]string{"RestartPolicy", string(data)}, " ")
}
