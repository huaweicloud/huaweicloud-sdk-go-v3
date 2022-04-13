package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateHealthMonitorRequest struct {
	Body *CreateHealthMonitorRequestBody `json:"body,omitempty"`
}

func (o CreateHealthMonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHealthMonitorRequest struct{}"
	}

	return strings.Join([]string{"CreateHealthMonitorRequest", string(data)}, " ")
}
