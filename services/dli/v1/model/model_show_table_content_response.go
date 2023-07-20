package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTableContentResponse Response Object
type ShowTableContentResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 表的列名称和类型。
	Schema *[]interface{} `json:"schema,omitempty"`

	// 预览的表内容。
	Rows           *[]interface{} `json:"rows,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowTableContentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTableContentResponse struct{}"
	}

	return strings.Join([]string{"ShowTableContentResponse", string(data)}, " ")
}
