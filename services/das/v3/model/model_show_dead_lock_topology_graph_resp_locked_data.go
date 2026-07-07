package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowDeadLockTopologyGraphRespLockedData struct {

	// 字段序号
	FieldIndex int64 `json:"field_index"`

	// 十六进制原始值
	Hex string `json:"hex"`

	// 可读解码值
	Decoded string `json:"decoded"`

	// 列名
	ColumnName string `json:"column_name"`
}

func (o ShowDeadLockTopologyGraphRespLockedData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeadLockTopologyGraphRespLockedData struct{}"
	}

	return strings.Join([]string{"ShowDeadLockTopologyGraphRespLockedData", string(data)}, " ")
}
