package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListNodeLimitSqlModelResponseResult struct {

	// 限流任务SQL_ID。
	SqlId *string `json:"sql_id,omitempty"`

	// 限流任务SQL模板。
	SqlModel *string `json:"sql_model,omitempty"`
}

func (o ListNodeLimitSqlModelResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodeLimitSqlModelResponseResult struct{}"
	}

	return strings.Join([]string{"ListNodeLimitSqlModelResponseResult", string(data)}, " ")
}
