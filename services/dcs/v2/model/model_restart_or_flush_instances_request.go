package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartOrFlushInstancesRequest Request Object
type RestartOrFlushInstancesRequest struct {
	Body *ChangeInstanceStatusBody `json:"body,omitempty"`
}

func (o RestartOrFlushInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartOrFlushInstancesRequest struct{}"
	}

	return strings.Join([]string{"RestartOrFlushInstancesRequest", string(data)}, " ")
}
