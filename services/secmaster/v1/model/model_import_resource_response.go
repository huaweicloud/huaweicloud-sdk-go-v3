package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportResourceResponse Response Object
type ImportResourceResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data           *ImportResourceResponseBodyData `json:"data,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ImportResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportResourceResponse struct{}"
	}

	return strings.Join([]string{"ImportResourceResponse", string(data)}, " ")
}
