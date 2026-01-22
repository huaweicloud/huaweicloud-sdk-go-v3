package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteServiceSetsRequest Request Object
type BatchDeleteServiceSetsRequest struct {
	Body *BatchDeleteServiceSetsDto `json:"body,omitempty"`
}

func (o BatchDeleteServiceSetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteServiceSetsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteServiceSetsRequest", string(data)}, " ")
}
