package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowGraphResponse struct {
	Graph *Graph1 `json:"graph,omitempty" xml:"graph"`

	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误信息。
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage"`

	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误码。
	ErrorCode      *string `json:"errorCode,omitempty" xml:"errorCode"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowGraphResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGraphResponse struct{}"
	}

	return strings.Join([]string{"ShowGraphResponse", string(data)}, " ")
}
