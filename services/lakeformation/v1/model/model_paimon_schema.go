package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PaimonSchema PaimonSchema模型。
type PaimonSchema struct {

	// 字段列表，定义表的所有列及其类型。
	Fields []PaimonField `json:"fields"`

	// 分区建集合
	PartitionKeys *[]string `json:"partition_keys,omitempty"`

	// 主键集合
	PrimaryKeys *[]string `json:"primary_keys,omitempty"`

	// Paimon表属性
	Options map[string]string `json:"options,omitempty"`
}

func (o PaimonSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PaimonSchema struct{}"
	}

	return strings.Join([]string{"PaimonSchema", string(data)}, " ")
}
