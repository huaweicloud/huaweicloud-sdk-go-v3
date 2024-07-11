package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAppRequest Request Object
type BatchDeleteAppRequest struct {
	Body *AppBatchDeleteRequest `json:"body,omitempty"`
}

func (o BatchDeleteAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAppRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteAppRequest", string(data)}, " ")
}
