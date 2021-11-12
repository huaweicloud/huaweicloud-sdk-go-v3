package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListGraphsResponse struct {
	// 图总个数。请求失败时为空。

	GraphCount *int32 `json:"graphCount,omitempty"`
	// 图列表。请求失败时为空。

	Graphs *[]Graph1 `json:"graphs,omitempty"`
	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误信息。

	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误码。

	ErrorCode      *string `json:"errorCode,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListGraphsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGraphsResponse struct{}"
	}

	return strings.Join([]string{"ListGraphsResponse", string(data)}, " ")
}
