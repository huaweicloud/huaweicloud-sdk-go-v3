package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeResourceResponse Response Object
type ChangeResourceResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data           *ChangeResourceRequestBody `json:"data,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ChangeResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeResourceResponse struct{}"
	}

	return strings.Join([]string{"ChangeResourceResponse", string(data)}, " ")
}
