package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Policy struct {

	// 备份周期的crontab表达式
	Period string `json:"period"`

	// 保留时长（天）
	RetentionDays int32 `json:"retention_days"`
}

func (o Policy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Policy struct{}"
	}

	return strings.Join([]string{"Policy", string(data)}, " ")
}
