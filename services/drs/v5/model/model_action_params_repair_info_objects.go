package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ActionParamsRepairInfoObjects struct {

	// 数据库名称。
	Database *string `json:"database,omitempty"`

	// schema名称。
	Schema *string `json:"schema,omitempty"`

	// 表名称。
	Table *string `json:"table,omitempty"`
}

func (o ActionParamsRepairInfoObjects) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionParamsRepairInfoObjects struct{}"
	}

	return strings.Join([]string{"ActionParamsRepairInfoObjects", string(data)}, " ")
}
