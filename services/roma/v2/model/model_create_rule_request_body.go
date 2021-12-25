package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRuleRequestBody struct {
	// 规则名称，支持英文大小写，数字，下划线和中划线,长度1-64

	Name string `json:"name"`
	// 应用ID

	AppId string `json:"app_id"`
	// 描述，长度0-200

	Description *string `json:"description,omitempty"`
	// 规则状态 0-启用 1-停用，不填写时默认为0-启用

	Status *int32 `json:"status,omitempty"`
	// 数据解析状态，0-启用 1-停用，不填写时默认为1-禁用

	DataParsingStatus *int32 `json:"data_parsing_status,omitempty"`
}

func (o CreateRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRuleRequestBody", string(data)}, " ")
}
