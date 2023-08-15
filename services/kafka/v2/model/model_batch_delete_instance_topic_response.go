package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteInstanceTopicResponse Response Object
type BatchDeleteInstanceTopicResponse struct {

	// Topic列表。
	Topics         *[]BatchDeleteInstanceTopicRespTopics `json:"topics,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o BatchDeleteInstanceTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstanceTopicResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstanceTopicResponse", string(data)}, " ")
}
