package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteInstanceTopicResponse Response Object
type BatchDeleteInstanceTopicResponse struct {

	// **参数解释**： Topic列表。
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
