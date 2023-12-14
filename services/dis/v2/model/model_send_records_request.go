package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendRecordsRequest Request Object
type SendRecordsRequest struct {
	Body *PutRecordsRequest `json:"body,omitempty"`
}

func (o SendRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendRecordsRequest struct{}"
	}

	return strings.Join([]string{"SendRecordsRequest", string(data)}, " ")
}
