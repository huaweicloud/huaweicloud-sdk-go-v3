package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWdrReportRequest Request Object
type CreateWdrReportRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	Body *CreateWdrReportRequestBody `json:"body,omitempty"`
}

func (o CreateWdrReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWdrReportRequest struct{}"
	}

	return strings.Join([]string{"CreateWdrReportRequest", string(data)}, " ")
}
