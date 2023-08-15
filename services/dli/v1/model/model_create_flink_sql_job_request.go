package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlinkSqlJobRequest Request Object
type CreateFlinkSqlJobRequest struct {
	Body *CreateFlinkSqlJobRequestBody `json:"body,omitempty"`
}

func (o CreateFlinkSqlJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkSqlJobRequest struct{}"
	}

	return strings.Join([]string{"CreateFlinkSqlJobRequest", string(data)}, " ")
}
