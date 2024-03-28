package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFlinkSqlJobRequest Request Object
type UpdateFlinkSqlJobRequest struct {

	// 作业ID。
	JobId int64 `json:"job_id"`

	Body *UpdateFlinkSqlJobRequestBody `json:"body,omitempty"`
}

func (o UpdateFlinkSqlJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlinkSqlJobRequest struct{}"
	}

	return strings.Join([]string{"UpdateFlinkSqlJobRequest", string(data)}, " ")
}
