package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatabaseForListTableResult 数据库表信息。
type DatabaseForListTableResult struct {

	// **参数解释**： 表名称。 **取值范围**： 不涉及。
	TableName string `json:"table_name"`
}

func (o DatabaseForListTableResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseForListTableResult struct{}"
	}

	return strings.Join([]string{"DatabaseForListTableResult", string(data)}, " ")
}
