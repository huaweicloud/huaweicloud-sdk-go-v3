package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCollectorChannelResponse Response Object
type UpdateCollectorChannelResponse struct {

	// 创建成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o UpdateCollectorChannelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCollectorChannelResponse struct{}"
	}

	return strings.Join([]string{"UpdateCollectorChannelResponse", string(data)}, " ")
}
