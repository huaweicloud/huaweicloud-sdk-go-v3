package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type PutRecordsRequest struct {
	Body *PutRecordsRequestBody `json:"body,omitempty"`
}

func (o PutRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutRecordsRequest struct{}"
	}

	return strings.Join([]string{"PutRecordsRequest", string(data)}, " ")
}
