package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFlinkSqlRequest Request Object
type UpdateFlinkSqlRequest struct {

	// 作业ID。
	JobId int64 `json:"job_id"`

	Body *UpdateFlinkSqlRequestBody `json:"body,omitempty"`
}

func (o UpdateFlinkSqlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlinkSqlRequest struct{}"
	}

	return strings.Join([]string{"UpdateFlinkSqlRequest", string(data)}, " ")
}
