package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExportTopSqlTemplatesDetailsResponse struct {

	// SQL模板列表。
	TopSqlTemplates *[]TopSqlTemplate `json:"top_sql_templates,omitempty"`

	// SQL模板总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ExportTopSqlTemplatesDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportTopSqlTemplatesDetailsResponse struct{}"
	}

	return strings.Join([]string{"ExportTopSqlTemplatesDetailsResponse", string(data)}, " ")
}
