package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIefSystemEventsResponse Response Object
type CreateIefSystemEventsResponse struct {

	// 请求是否成功
	IsSuccess *bool `json:"is_success,omitempty"`

	// 返回信息
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateIefSystemEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIefSystemEventsResponse struct{}"
	}

	return strings.Join([]string{"CreateIefSystemEventsResponse", string(data)}, " ")
}
