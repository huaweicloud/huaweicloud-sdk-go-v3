package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSqlJobTemplateRequest Request Object
type UpdateSqlJobTemplateRequest struct {
	SqlId string `json:"sql_id"`

	Body *UpdateSqlJobTemplateRequestBody `json:"body,omitempty"`
}

func (o UpdateSqlJobTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSqlJobTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateSqlJobTemplateRequest", string(data)}, " ")
}
