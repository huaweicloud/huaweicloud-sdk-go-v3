package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
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
