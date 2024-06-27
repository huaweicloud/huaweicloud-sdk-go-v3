package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DepParamInstance struct {

	// 比较值
	CompareValue *string `json:"compareValue,omitempty"`

	// 比较
	Comparison *string `json:"comparison,omitempty"`

	// 条件
	Condition *bool `json:"condition,omitempty"`

	// 无效值列表
	InValidValues *[]string `json:"inValidValues,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 空值信息
	NullInfo *string `json:"nullInfo,omitempty"`

	// 关系映射，key为整数，value为字符串
	RelationMap map[string]string `json:"relationMap,omitempty"`

	// 关系数量
	RelationNum *int32 `json:"relationNum,omitempty"`

	// 值列表
	Values *[]string `json:"values,omitempty"`
}

func (o DepParamInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DepParamInstance struct{}"
	}

	return strings.Join([]string{"DepParamInstance", string(data)}, " ")
}
