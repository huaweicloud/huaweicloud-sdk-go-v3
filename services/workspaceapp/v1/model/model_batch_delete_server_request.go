package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteServerRequest Request Object
type BatchDeleteServerRequest struct {
	Body *BatchDeleteServerReq `json:"body,omitempty"`
}

func (o BatchDeleteServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteServerRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteServerRequest", string(data)}, " ")
}
