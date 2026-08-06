package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteConnectionNewRequest Request Object
type BatchDeleteConnectionNewRequest struct {
	Body *BatchDeleteConnectionNewRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteConnectionNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteConnectionNewRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteConnectionNewRequest", string(data)}, " ")
}
