package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportSlowSqlTemplatesDetailsResponse Response Object
type ExportSlowSqlTemplatesDetailsResponse struct {

	// 慢日志模板数据列表。
	SlowSqlTemplates *[]SlowSqlTemplate `json:"slow_sql_templates,omitempty"`

	// 慢日志模板总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ExportSlowSqlTemplatesDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSlowSqlTemplatesDetailsResponse struct{}"
	}

	return strings.Join([]string{"ExportSlowSqlTemplatesDetailsResponse", string(data)}, " ")
}
