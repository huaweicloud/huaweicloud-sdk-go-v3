package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLogStreamResponse Response Object
type DeleteLogStreamResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteLogStreamResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLogStreamResponse struct{}"
	}

	return strings.Join([]string{"DeleteLogStreamResponse", string(data)}, " ")
}
