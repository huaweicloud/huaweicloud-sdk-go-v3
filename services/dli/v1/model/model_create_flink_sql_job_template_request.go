package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlinkSqlJobTemplateRequest Request Object
type CreateFlinkSqlJobTemplateRequest struct {
	Body *CreateFlinkSqlJobTemplateRequestBody `json:"body,omitempty"`
}

func (o CreateFlinkSqlJobTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkSqlJobTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateFlinkSqlJobTemplateRequest", string(data)}, " ")
}
