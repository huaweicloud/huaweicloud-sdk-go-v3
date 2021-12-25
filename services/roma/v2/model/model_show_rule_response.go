package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRuleResponse struct {
	// 权限

	Permissions *[]string `json:"permissions,omitempty"`
	// 规则ID

	RuleId *int32 `json:"rule_id,omitempty"`
	// 规则名称，支持英文大小写，数字，下划线和中划线,长度1-64

	Name *string `json:"name,omitempty"`
	// 应用ID

	AppId *string `json:"app_id,omitempty"`
	// 应用名称

	AppName *string `json:"app_name,omitempty"`
	// 描述

	Description *string `json:"description,omitempty"`
	// 规则状态 0-启用 1-停用

	Status *int32 `json:"status,omitempty"`
	// 数据解析状态，ENABLE时data_parsing必填 0-启用 1-停用

	DataParsingStatus *int32 `json:"data_parsing_status,omitempty"`
	// SQL查询字段

	SqlField *string `json:"sql_field,omitempty"`
	// SQL查询条件

	SqlWhere *string `json:"sql_where,omitempty"`
	// 完整的规则表达式

	RuleExpress *string `json:"rule_express,omitempty"`

	CreatedUser *CreatedUser `json:"created_user,omitempty"`

	LastUpdatedUser *LastUpdatedUser `json:"last_updated_user,omitempty"`
	// 创建时间，timestamp(ms)，使用UTC时区

	CreatedDatetime *int64 `json:"created_datetime,omitempty"`
	// 最后修改时间，timestamp(ms)，使用UTC时区

	LastUpdatedDatetime *int64 `json:"last_updated_datetime,omitempty"`
	HttpStatusCode      int    `json:"-"`
}

func (o ShowRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowRuleResponse", string(data)}, " ")
}
