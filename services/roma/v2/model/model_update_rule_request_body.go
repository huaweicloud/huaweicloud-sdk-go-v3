package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRuleRequestBody struct {
	// 规则名称，支持英文大小写，数字，下划线和中划线,长度1-64

	Name *string `json:"name,omitempty"`
	// 描述，长度0-200

	Description *string `json:"description,omitempty"`
	// 规则状态 0-启用 1-停用，不填写时默认为0-启用

	Status *int32 `json:"status,omitempty"`
	// 数据解析状态，0-启用 1-停用，不填写时默认为1-禁用

	DataParsingStatus *int32 `json:"data_parsing_status,omitempty"`
	// SQL查询字段

	SqlField *string `json:"sql_field,omitempty"`
	// SQL查询条件

	SqlWhere *string `json:"sql_where,omitempty"`
}

func (o UpdateRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateRuleRequestBody", string(data)}, " ")
}
