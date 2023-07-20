package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIefMessageChannelResponse Response Object
type CreateIefMessageChannelResponse struct {

	// 请求是否成功
	IsSuccess *bool `json:"is_success,omitempty"`

	// 返回信息
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateIefMessageChannelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIefMessageChannelResponse struct{}"
	}

	return strings.Join([]string{"CreateIefMessageChannelResponse", string(data)}, " ")
}
