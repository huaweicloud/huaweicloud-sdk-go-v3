package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScattersUsingRequest Request Object
type ListScattersUsingRequest struct {

	// 服务id
	ServiceId string `json:"service_id"`

	Body *DashboardDataQuery `json:"body,omitempty"`
}

func (o ListScattersUsingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScattersUsingRequest struct{}"
	}

	return strings.Join([]string{"ListScattersUsingRequest", string(data)}, " ")
}
