package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckHealthReportTaskResponse Response Object
type CheckHealthReportTaskResponse struct {

	// 是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CheckHealthReportTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckHealthReportTaskResponse struct{}"
	}

	return strings.Join([]string{"CheckHealthReportTaskResponse", string(data)}, " ")
}
