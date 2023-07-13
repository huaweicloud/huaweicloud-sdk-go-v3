package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePublicIpRequest Request Object
type BatchDeletePublicIpRequest struct {
	Body *BatchDeletePublicIpRequestBody `json:"body,omitempty"`
}

func (o BatchDeletePublicIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePublicIpRequest struct{}"
	}

	return strings.Join([]string{"BatchDeletePublicIpRequest", string(data)}, " ")
}
