package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogConfigDto 日志配置
type LogConfigDto struct {

	// 应用日志文件大小限制，单位MB，默认50，取值范围10-1000
	Size *int32 `json:"size,omitempty"`

	// 应用日志级别，可选项：on/off/trace/debug/info/warn/error/fatal，当type为LTS时有效。当选择非on/off的选项时，将只采集大于等于指定级别的日志。
	Level *string `json:"level,omitempty"`

	// 应用日志rotate个数，默认5，取值范围1-10
	RotateNum *int32 `json:"rotate_num,omitempty"`

	// 应用日志rotate周期，可选项： daily/monthly/weekly/yearly
	RotatePeriod *string `json:"rotate_period,omitempty"`

	// LTS:将日志发送到LTS, local 本地日志
	Type *string `json:"type,omitempty"`

	// app:部署到边缘设备上的应用的日志, system 边缘设备上系统的日志
	Component *string `json:"component,omitempty"`
}

func (o LogConfigDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogConfigDto struct{}"
	}

	return strings.Join([]string{"LogConfigDto", string(data)}, " ")
}
