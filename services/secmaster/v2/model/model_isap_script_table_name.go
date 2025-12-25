package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapScriptTableName 表名
type IsapScriptTableName struct {
}

func (o IsapScriptTableName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapScriptTableName struct{}"
	}

	return strings.Join([]string{"IsapScriptTableName", string(data)}, " ")
}
