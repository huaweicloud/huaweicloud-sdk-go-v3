package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDeleteLabelRequest struct {
	Body *BatchDeleteLabelReq `json:"body,omitempty"`
}

func (o BatchDeleteLabelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteLabelRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteLabelRequest", string(data)}, " ")
}
