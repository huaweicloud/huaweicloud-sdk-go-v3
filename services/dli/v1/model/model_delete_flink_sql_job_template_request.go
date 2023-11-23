package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFlinkSqlJobTemplateRequest Request Object
type DeleteFlinkSqlJobTemplateRequest struct {

	// 模板ID。
	TemplateId int64 `json:"template_id"`
}

func (o DeleteFlinkSqlJobTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFlinkSqlJobTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteFlinkSqlJobTemplateRequest", string(data)}, " ")
}
