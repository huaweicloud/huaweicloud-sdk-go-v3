package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStartTransferTaskResponse Response Object
type BatchStartTransferTaskResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchStartTransferTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartTransferTaskResponse struct{}"
	}

	return strings.Join([]string{"BatchStartTransferTaskResponse", string(data)}, " ")
}
