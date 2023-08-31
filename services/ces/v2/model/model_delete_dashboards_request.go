package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDashboardsRequest Request Object
type DeleteDashboardsRequest struct {
	Body *BatchDeleteDashboardRequestBody `json:"body,omitempty"`
}

func (o DeleteDashboardsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDashboardsRequest struct{}"
	}

	return strings.Join([]string{"DeleteDashboardsRequest", string(data)}, " ")
}
