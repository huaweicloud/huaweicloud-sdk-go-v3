package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlinkSqlJobGraphRequest Request Object
type CreateFlinkSqlJobGraphRequest struct {
	JobId string `json:"job_id"`

	Body *CreateFlinkSqlJobGraphRequestBody `json:"body,omitempty"`
}

func (o CreateFlinkSqlJobGraphRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkSqlJobGraphRequest struct{}"
	}

	return strings.Join([]string{"CreateFlinkSqlJobGraphRequest", string(data)}, " ")
}
