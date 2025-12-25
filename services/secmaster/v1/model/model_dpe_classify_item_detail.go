package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DpeClassifyItemDetail struct {

	// 映射id
	Id *string `json:"id,omitempty"`

	// 映射id
	ClassifierId *string `json:"classifier_id,omitempty"`

	// 映射id
	ClassifierTypeId *string `json:"classifier_type_id,omitempty"`

	// 映射id
	MappingId *string `json:"mapping_id,omitempty"`

	// 分类优先级
	ClassifierOrder *int32 `json:"classifier_order,omitempty"`

	// 表达式
	Expression *string `json:"expression,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o DpeClassifyItemDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DpeClassifyItemDetail struct{}"
	}

	return strings.Join([]string{"DpeClassifyItemDetail", string(data)}, " ")
}
