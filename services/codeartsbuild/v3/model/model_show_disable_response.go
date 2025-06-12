package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDisableResponse Response Object
type ShowDisableResponse struct {
	Result *IsDisableResult `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDisableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDisableResponse struct{}"
	}

	return strings.Join([]string{"ShowDisableResponse", string(data)}, " ")
}
