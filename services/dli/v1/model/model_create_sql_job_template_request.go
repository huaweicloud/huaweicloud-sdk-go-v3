package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSqlJobTemplateRequest Request Object
type CreateSqlJobTemplateRequest struct {
	Body *CreateSqlJobTemplateRequestBody `json:"body,omitempty"`
}

func (o CreateSqlJobTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlJobTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateSqlJobTemplateRequest", string(data)}, " ")
}
