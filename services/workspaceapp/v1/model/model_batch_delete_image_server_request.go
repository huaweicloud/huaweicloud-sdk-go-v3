package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteImageServerRequest Request Object
type BatchDeleteImageServerRequest struct {
	Body *BatchDeleteImageServerReq `json:"body,omitempty"`
}

func (o BatchDeleteImageServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteImageServerRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteImageServerRequest", string(data)}, " ")
}
