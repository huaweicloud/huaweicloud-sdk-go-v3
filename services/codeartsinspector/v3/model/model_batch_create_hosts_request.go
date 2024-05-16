package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateHostsRequest Request Object
type BatchCreateHostsRequest struct {
	Body *BatchCreateHostsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateHostsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateHostsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateHostsRequest", string(data)}, " ")
}
