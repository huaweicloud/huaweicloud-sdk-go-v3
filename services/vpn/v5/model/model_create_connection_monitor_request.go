package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectionMonitorRequest Request Object
type CreateConnectionMonitorRequest struct {
	Body *CreateConnectionMonitorRequestBody `json:"body,omitempty"`
}

func (o CreateConnectionMonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionMonitorRequest struct{}"
	}

	return strings.Join([]string{"CreateConnectionMonitorRequest", string(data)}, " ")
}
