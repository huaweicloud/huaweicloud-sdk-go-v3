package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceHealthReportTaskNewResponse Response Object
type CreateInstanceHealthReportTaskNewResponse struct {

	// 是否成功
	CreateSuccess  *bool `json:"create_success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreateInstanceHealthReportTaskNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceHealthReportTaskNewResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceHealthReportTaskNewResponse", string(data)}, " ")
}
