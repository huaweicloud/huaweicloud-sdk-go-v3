package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKafkaScalePreCheckInfoResponse Response Object
type ShowKafkaScalePreCheckInfoResponse struct {
	Body           *[]ShowKafkaScalePreCheckInfoResponseBody `json:"body,omitempty"`
	HttpStatusCode int                                       `json:"-"`
}

func (o ShowKafkaScalePreCheckInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaScalePreCheckInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowKafkaScalePreCheckInfoResponse", string(data)}, " ")
}
