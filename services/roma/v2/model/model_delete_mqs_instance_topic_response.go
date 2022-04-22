package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteMqsInstanceTopicResponse struct {

	// Topic列表。
	Topics         *[]DeleteMqsInstanceTopicRespTopics `json:"topics,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o DeleteMqsInstanceTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMqsInstanceTopicResponse struct{}"
	}

	return strings.Join([]string{"DeleteMqsInstanceTopicResponse", string(data)}, " ")
}
