package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlinkJobExecutionGraphRequest Request Object
type ShowFlinkJobExecutionGraphRequest struct {

	// 作业ID。
	JobId int64 `json:"job_id"`
}

func (o ShowFlinkJobExecutionGraphRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkJobExecutionGraphRequest struct{}"
	}

	return strings.Join([]string{"ShowFlinkJobExecutionGraphRequest", string(data)}, " ")
}
