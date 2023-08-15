package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabasesResponse Response Object
type ListDatabasesResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 数据库的总数。
	DatabaseCount *int32 `json:"database_count,omitempty"`

	// 查询所有数据库的响应参数。
	Databases      *[]Database `json:"databases,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListDatabasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabasesResponse struct{}"
	}

	return strings.Join([]string{"ListDatabasesResponse", string(data)}, " ")
}
