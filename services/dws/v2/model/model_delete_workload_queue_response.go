package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteWorkloadQueueResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteWorkloadQueueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkloadQueueResponse struct{}"
	}

	return strings.Join([]string{"DeleteWorkloadQueueResponse", string(data)}, " ")
}
