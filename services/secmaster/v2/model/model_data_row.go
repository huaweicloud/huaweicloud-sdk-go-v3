package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataRow struct {

	// 数据类型
	Kind *string `json:"kind,omitempty"`

	// 数据行
	Fields *[]interface{} `json:"fields,omitempty"`
}

func (o DataRow) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataRow struct{}"
	}

	return strings.Join([]string{"DataRow", string(data)}, " ")
}
