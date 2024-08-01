package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConsumerGroupHeartBeatResponse Response Object
type ConsumerGroupHeartBeatResponse struct {
	Body           *[]string `json:"body,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ConsumerGroupHeartBeatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumerGroupHeartBeatResponse struct{}"
	}

	return strings.Join([]string{"ConsumerGroupHeartBeatResponse", string(data)}, " ")
}
