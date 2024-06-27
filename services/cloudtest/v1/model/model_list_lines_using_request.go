package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLinesUsingRequest Request Object
type ListLinesUsingRequest struct {

	// 服务id
	ServiceId string `json:"service_id"`

	Body *DashboardDataQuery `json:"body,omitempty"`
}

func (o ListLinesUsingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLinesUsingRequest struct{}"
	}

	return strings.Join([]string{"ListLinesUsingRequest", string(data)}, " ")
}
