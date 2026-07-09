package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogRetention 修改审计日志保存时间信息bean
type UpdateLogRetention struct {

	// 设定的日志保存时间信息,正整数。
	RetentionDays *int32 `json:"retention_days,omitempty"`
}

func (o UpdateLogRetention) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogRetention struct{}"
	}

	return strings.Join([]string{"UpdateLogRetention", string(data)}, " ")
}
