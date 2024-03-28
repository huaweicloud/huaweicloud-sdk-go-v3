package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatasourceConnectionsResponse Response Object
type ListDatasourceConnectionsResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 跨源连接信息列表。
	Connections *[]DatasourceConnection `json:"connections,omitempty"`

	// 总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDatasourceConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatasourceConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ListDatasourceConnectionsResponse", string(data)}, " ")
}
