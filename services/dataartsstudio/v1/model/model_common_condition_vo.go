package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommonConditionVo struct {

	// 名称
	Name string `json:"name"`

	// 字段id信息， 格式：table_id.field_id
	FieldIds []string `json:"field_ids"`

	// 字段名称信息， 格式：表名称.字段名称
	FieldNames *[]string `json:"field_names,omitempty"`

	// 计算表达式
	CalExp string `json:"cal_exp"`

	// 计算表达式id
	CalFnIds []int64 `json:"cal_fn_ids"`

	// 前端表达式配置，用于前端数据恢复
	FrontConfigs *string `json:"front_configs,omitempty"`

	// id
	Id *int64 `json:"id,omitempty"`
}

func (o CommonConditionVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonConditionVo struct{}"
	}

	return strings.Join([]string{"CommonConditionVo", string(data)}, " ")
}
