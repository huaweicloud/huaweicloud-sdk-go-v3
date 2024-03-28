package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkSqlJobTemplateId 删除模板信息。
type FlinkSqlJobTemplateId struct {

	// 模板ID。
	TemplateId *int64 `json:"template_id,omitempty"`
}

func (o FlinkSqlJobTemplateId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkSqlJobTemplateId struct{}"
	}

	return strings.Join([]string{"FlinkSqlJobTemplateId", string(data)}, " ")
}
