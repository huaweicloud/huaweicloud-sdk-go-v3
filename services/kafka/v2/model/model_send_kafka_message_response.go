package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendKafkaMessageResponse Response Object
type SendKafkaMessageResponse struct {
	Topic *string `json:"topic,omitempty"`

	Body *string `json:"body,omitempty"`

	PropertyList   *[]interface{} `json:"property_list,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o SendKafkaMessageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendKafkaMessageResponse struct{}"
	}

	return strings.Join([]string{"SendKafkaMessageResponse", string(data)}, " ")
}
