package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOneDashboardRequest Request Object
type CreateOneDashboardRequest struct {
	Body *CreateDashboardRequestBody `json:"body,omitempty"`
}

func (o CreateOneDashboardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOneDashboardRequest struct{}"
	}

	return strings.Join([]string{"CreateOneDashboardRequest", string(data)}, " ")
}
