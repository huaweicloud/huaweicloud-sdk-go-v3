package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotifySmnTopicConfigResponse Response Object
type NotifySmnTopicConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NotifySmnTopicConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotifySmnTopicConfigResponse struct{}"
	}

	return strings.Join([]string{"NotifySmnTopicConfigResponse", string(data)}, " ")
}
