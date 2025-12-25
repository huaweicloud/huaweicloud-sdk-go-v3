package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SqlValiTableName Table name 表名
type SqlValiTableName struct {
}

func (o SqlValiTableName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlValiTableName struct{}"
	}

	return strings.Join([]string{"SqlValiTableName", string(data)}, " ")
}
