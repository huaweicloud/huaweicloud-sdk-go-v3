package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDataChannelResponse Response Object
type DeleteDataChannelResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDataChannelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataChannelResponse struct{}"
	}

	return strings.Join([]string{"DeleteDataChannelResponse", string(data)}, " ")
}
