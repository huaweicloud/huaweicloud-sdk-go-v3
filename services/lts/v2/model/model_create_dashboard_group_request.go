package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDashboardGroupRequest Request Object
type CreateDashboardGroupRequest struct {
	Body *CreateDashboardGroupReq `json:"body,omitempty"`
}

func (o CreateDashboardGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDashboardGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateDashboardGroupRequest", string(data)}, " ")
}
