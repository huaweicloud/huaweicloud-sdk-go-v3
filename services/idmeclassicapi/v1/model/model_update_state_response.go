package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStateResponse Response Object
type UpdateStateResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]interface{} `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdateStateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStateResponse struct{}"
	}

	return strings.Join([]string{"UpdateStateResponse", string(data)}, " ")
}
