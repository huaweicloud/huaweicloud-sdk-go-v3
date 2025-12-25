package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCollectorChannelResponse Response Object
type CreateCollectorChannelResponse struct {

	// 创建成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreateCollectorChannelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCollectorChannelResponse struct{}"
	}

	return strings.Join([]string{"CreateCollectorChannelResponse", string(data)}, " ")
}
