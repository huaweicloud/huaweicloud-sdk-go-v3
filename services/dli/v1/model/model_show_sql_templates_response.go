package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlTemplatesResponse Response Object
type ShowSqlTemplatesResponse struct {

	// 是否成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。执行失败时，用于显示执行失败的原因。
	Message *string `json:"message,omitempty"`

	// SQL模板总数。
	SqlCount *int32 `json:"sql_count,omitempty"`

	Sqls           *[]SqlsResp `json:"sqls,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowSqlTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlTemplatesResponse", string(data)}, " ")
}
