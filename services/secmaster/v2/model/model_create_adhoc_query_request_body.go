package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAdhocQueryRequestBody struct {

	// UUID
	SessionId *string `json:"session_id,omitempty"`

	// 具体查询
	Query *string `json:"query,omitempty"`

	// 查询类型
	QueryType *string `json:"query_type,omitempty"`

	// 起始范围
	From *int64 `json:"from,omitempty"`

	// 终止范围
	To *int64 `json:"to,omitempty"`

	// 脚本参数列表
	ScriptParams *[]ScriptParam `json:"script_params,omitempty"`

	// 时区
	TimeZone *string `json:"time_zone,omitempty"`

	// 起始范围
	SqlRenderFrom *int64 `json:"sql_render_from,omitempty"`

	// 终止范围
	SqlRenderTo *int64 `json:"sql_render_to,omitempty"`

	// 源页
	SourcePage *string `json:"source_page,omitempty"`
}

func (o CreateAdhocQueryRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAdhocQueryRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAdhocQueryRequestBody", string(data)}, " ")
}
