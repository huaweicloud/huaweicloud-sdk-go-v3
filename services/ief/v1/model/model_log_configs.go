package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogConfigs 边缘节点日志配置
type LogConfigs struct {

	// 数据库保存的主键，不需要关注此字段。
	Id *string `json:"id,omitempty"`

	// 应用日志文件大小限制，单位MB，默认50，取值范围10-1000。
	Size *int32 `json:"size,omitempty"`

	// - 当type为LTS时，应用级日志可配置为On或Off；系统级日志可配置为Off/Error/Warning/Info/Debug； - 当type为local时，无需配置level。
	Level *string `json:"level,omitempty"`

	// 应用日志rotate个数，默认5，取值范围1-10。
	RotateNum *int32 `json:"rotate_num,omitempty"`

	// 应用日志rotate周期，可选项： - daily - monthly - weekly - yearly
	RotatePeriod *string `json:"rotate_period,omitempty"`

	// - LTS：将日志发送到AOM - local：本地日志
	Type string `json:"type"`

	// - app：部署到边缘节点上的应用的日志 - system：边缘节点上系统的日志
	Component string `json:"component"`
}

func (o LogConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogConfigs struct{}"
	}

	return strings.Join([]string{"LogConfigs", string(data)}, " ")
}
