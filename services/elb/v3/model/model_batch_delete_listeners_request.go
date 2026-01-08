package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteListenersRequest Request Object
type BatchDeleteListenersRequest struct {
	Body *BatchDeleteListenersRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteListenersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteListenersRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteListenersRequest", string(data)}, " ")
}
