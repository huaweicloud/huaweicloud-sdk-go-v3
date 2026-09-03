package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InvokeWdrReportRequest Request Object
type InvokeWdrReportRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	Body *InvokeWdrReportRequestBody `json:"body,omitempty"`
}

func (o InvokeWdrReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvokeWdrReportRequest struct{}"
	}

	return strings.Join([]string{"InvokeWdrReportRequest", string(data)}, " ")
}
