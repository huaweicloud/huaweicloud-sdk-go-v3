package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePushChannelRequest Request Object
type CreatePushChannelRequest struct {
	Body *CreateChannelRequestDto `json:"body,omitempty"`
}

func (o CreatePushChannelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePushChannelRequest struct{}"
	}

	return strings.Join([]string{"CreatePushChannelRequest", string(data)}, " ")
}
