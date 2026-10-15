package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GeneralDataFilterResp 数据过滤对象响应体
type GeneralDataFilterResp struct {

	// 源数据库的过滤条件SQL表达式。
	SourceFilterCondition *string `json:"source_filter_condition,omitempty"`

	// 目标数据库的过滤条件SQL表达式。
	TargetFilterCondition *string `json:"target_filter_condition,omitempty"`

	// 标记源库和目标库过滤条件是否分开配置。
	ConditionSeparated *string `json:"condition_separated,omitempty"`
}

func (o GeneralDataFilterResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GeneralDataFilterResp struct{}"
	}

	return strings.Join([]string{"GeneralDataFilterResp", string(data)}, " ")
}
