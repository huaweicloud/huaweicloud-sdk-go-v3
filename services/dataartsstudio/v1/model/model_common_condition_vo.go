package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommonConditionVo struct {

	// 名称。
	Name string `json:"name"`

	// 字段id信息，格式：table_id.field_id。
	FieldIds []string `json:"field_ids"`

	// 字段名称信息，格式：表名称.字段名称。
	FieldNames *[]string `json:"field_names,omitempty"`

	// 计算表达式，形如'${table_id.column_id} > 1'，其中table_id表示引用字段所属表ID，column_id表示引用字段ID。
	CalExp string `json:"cal_exp"`

	// 引用函数ID，填写String类型替代Long类型。
	CalFnIds []string `json:"cal_fn_ids"`

	// 前端表达式配置，用于前端数据恢复。
	FrontConfigs *string `json:"front_configs,omitempty"`

	// 通用限定ID，只读，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`
}

func (o CommonConditionVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonConditionVo struct{}"
	}

	return strings.Join([]string{"CommonConditionVo", string(data)}, " ")
}
