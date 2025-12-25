package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapTableName 表名称
type IsapTableName struct {
}

func (o IsapTableName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapTableName struct{}"
	}

	return strings.Join([]string{"IsapTableName", string(data)}, " ")
}
