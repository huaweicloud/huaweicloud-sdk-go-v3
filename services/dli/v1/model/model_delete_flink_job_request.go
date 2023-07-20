package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFlinkJobRequest Request Object
type DeleteFlinkJobRequest struct {

	// 作业ID。
	JobId int64 `json:"job_id"`
}

func (o DeleteFlinkJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFlinkJobRequest struct{}"
	}

	return strings.Join([]string{"DeleteFlinkJobRequest", string(data)}, " ")
}
