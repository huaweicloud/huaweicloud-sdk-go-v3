package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlSampleTemplatesResponse Response Object
type ListSqlSampleTemplatesResponse struct {

	// 请求执行是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 样例模板信息。
	Sqls *[]SqlSampleTemplate `json:"sqls,omitempty"`

	// 样例模板个数。
	SqlCount       *int32 `json:"sqlCount,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSqlSampleTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlSampleTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListSqlSampleTemplatesResponse", string(data)}, " ")
}
