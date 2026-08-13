package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PaimonType Paimon的类型定义，支持简单类型和复杂类型，遵循Paimon 类型系统标准。
type PaimonType struct {

	// Paimon类型名称
	Name string `json:"name"`

	// 精度，适用于 DECIMAL(p,s) 的 p（1-38），以及 TIME(p)/TIMESTAMP(p)/TIMESTAMP_LTZ(p) 的小数秒精度 p（0-9）。
	Precision *int32 `json:"precision,omitempty"`

	// 标度，适用于 DECIMAL(p,s) 的 s（0-precision）。
	Scale *int32 `json:"scale,omitempty"`

	// CHAR(n)、VARCHAR(n)、BINARY(n)、VARBINARY(n)的长度（n）
	Length *int32 `json:"length,omitempty"`
}

func (o PaimonType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PaimonType struct{}"
	}

	return strings.Join([]string{"PaimonType", string(data)}, " ")
}
