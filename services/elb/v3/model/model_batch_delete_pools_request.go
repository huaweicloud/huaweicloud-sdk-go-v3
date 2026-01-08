package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePoolsRequest Request Object
type BatchDeletePoolsRequest struct {
	Body *BatchDeletePoolsRequestBody `json:"body,omitempty"`
}

func (o BatchDeletePoolsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePoolsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeletePoolsRequest", string(data)}, " ")
}
