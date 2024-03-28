package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFlinkSqlJobTemplateRequest Request Object
type UpdateFlinkSqlJobTemplateRequest struct {
	TemplateId int64 `json:"template_id"`

	Body *UpdateFlinkSqlJobTemplateRequestBody `json:"body,omitempty"`
}

func (o UpdateFlinkSqlJobTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlinkSqlJobTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateFlinkSqlJobTemplateRequest", string(data)}, " ")
}
