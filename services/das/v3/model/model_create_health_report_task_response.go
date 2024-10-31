package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHealthReportTaskResponse Response Object
type CreateHealthReportTaskResponse struct {

	// 诊断任务创建是否成功
	CreateSuccess  *bool `json:"create_success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreateHealthReportTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHealthReportTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateHealthReportTaskResponse", string(data)}, " ")
}
