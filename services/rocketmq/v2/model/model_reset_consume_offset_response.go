package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetConsumeOffsetResponse Response Object
type ResetConsumeOffsetResponse struct {

	// 重置的队列。
	Queues         *[]ResetConsumeOffsetRespQueues `json:"queues,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ResetConsumeOffsetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetConsumeOffsetResponse struct{}"
	}

	return strings.Join([]string{"ResetConsumeOffsetResponse", string(data)}, " ")
}
