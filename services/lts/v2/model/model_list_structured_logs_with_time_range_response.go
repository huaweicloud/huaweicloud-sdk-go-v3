package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStructuredLogsWithTimeRangeResponse Response Object
type ListStructuredLogsWithTimeRangeResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListStructuredLogsWithTimeRangeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStructuredLogsWithTimeRangeResponse struct{}"
	}

	return strings.Join([]string{"ListStructuredLogsWithTimeRangeResponse", string(data)}, " ")
}
