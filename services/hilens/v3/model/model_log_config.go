package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogConfig 日志配置列表
type LogConfig struct {

	// app：应用日志。 system：系统的日志
	Component string `json:"component"`

	// 系统级日志可配置为/error/warning/info/debug ; 不传会默认为info。
	Level *string `json:"level,omitempty"`

	// 日志rotate个数，默认5，hilens取值范围1-30，ief取值范围1-10
	RotateNum *int32 `json:"rotate_num,omitempty"`

	// 日志rotate周期，可选项，只支持ief：daily monthly weekly yearly.
	RotatePeriod *string `json:"rotate_period,omitempty"`

	// 应用日志文件大小限制，单位MB，默认50，取值范围10-1000。
	Size *int32 `json:"size,omitempty"`

	// - LTS 将日志发送到云日志服务（Log Tank Service，简称LTS） - local 本地日志
	Type string `json:"type"`
}

func (o LogConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogConfig struct{}"
	}

	return strings.Join([]string{"LogConfig", string(data)}, " ")
}
