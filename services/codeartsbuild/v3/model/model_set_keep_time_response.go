package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetKeepTimeResponse Response Object
type SetKeepTimeResponse struct {
	Result *SetKeepTimeResult `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetKeepTimeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetKeepTimeResponse struct{}"
	}

	return strings.Join([]string{"SetKeepTimeResponse", string(data)}, " ")
}
