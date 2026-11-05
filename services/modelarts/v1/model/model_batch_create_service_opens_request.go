package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateServiceOpensRequest Request Object
type BatchCreateServiceOpensRequest struct {
	Body *ServiceBatchOpenRequest `json:"body,omitempty"`
}

func (o BatchCreateServiceOpensRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateServiceOpensRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateServiceOpensRequest", string(data)}, " ")
}
