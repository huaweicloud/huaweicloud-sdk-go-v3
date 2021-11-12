package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateGraphResponse struct {
	// 图ID。

	Id *string `json:"id,omitempty"`
	// 图名称。

	Name *string `json:"name,omitempty"`
	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误信息。

	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误码。

	ErrorCode      *string `json:"errorCode,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateGraphResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGraphResponse struct{}"
	}

	return strings.Join([]string{"CreateGraphResponse", string(data)}, " ")
}
