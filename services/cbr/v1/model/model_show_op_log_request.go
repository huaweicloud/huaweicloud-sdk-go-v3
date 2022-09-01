package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowOpLogRequest struct {

	// 任务ID
	OperationLogId string `json:"operation_log_id" xml:"operation_log_id"`
}

func (o ShowOpLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOpLogRequest struct{}"
	}

	return strings.Join([]string{"ShowOpLogRequest", string(data)}, " ")
}
