package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateChannelRequest struct {
	Body *ChannelCreateReq `json:"body,omitempty" xml:"body"`
}

func (o CreateChannelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateChannelRequest struct{}"
	}

	return strings.Join([]string{"CreateChannelRequest", string(data)}, " ")
}
