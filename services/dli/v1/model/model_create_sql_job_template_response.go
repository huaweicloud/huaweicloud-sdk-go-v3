package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSqlJobTemplateResponse Response Object
type CreateSqlJobTemplateResponse struct {

	// 是否成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。执行失败时，用于显示执行失败的原因。
	Message *string `json:"message,omitempty"`

	// 新增SQL模板的ID。
	SqlId *string `json:"sql_id,omitempty"`

	// SQL模板分组名称。
	Group          *string `json:"group,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSqlJobTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlJobTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateSqlJobTemplateResponse", string(data)}, " ")
}
