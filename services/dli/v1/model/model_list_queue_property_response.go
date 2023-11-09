package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueuePropertyResponse Response Object
type ListQueuePropertyResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message *string `json:"message,omitempty"`

	Properties     *[]ListQueuePropertyRespProperties `json:"properties,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ListQueuePropertyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueuePropertyResponse struct{}"
	}

	return strings.Join([]string{"ListQueuePropertyResponse", string(data)}, " ")
}
