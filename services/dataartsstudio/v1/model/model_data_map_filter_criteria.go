package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataMapFilterCriteria 过滤条件
type DataMapFilterCriteria struct {

	// 属性
	Attribute *string `json:"attribute,omitempty"`

	// 操作表示，默认值EQ
	Operator *string `json:"operator,omitempty"`

	// 值
	Value *interface{} `json:"value,omitempty"`

	Condition *ConditionInfo `json:"condition,omitempty"`

	// 条件准则
	Criterion *[]DataMapFilterCriteria `json:"criterion,omitempty"`
}

func (o DataMapFilterCriteria) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataMapFilterCriteria struct{}"
	}

	return strings.Join([]string{"DataMapFilterCriteria", string(data)}, " ")
}
