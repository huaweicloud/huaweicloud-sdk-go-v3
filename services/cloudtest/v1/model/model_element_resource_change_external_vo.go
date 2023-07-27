package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ElementResourceChangeExternalVo 数据信息
type ElementResourceChangeExternalVo struct {

	// 工作项类型
	TrackerName *string `json:"tracker_name,omitempty"`

	// 值
	Value *string `json:"value,omitempty"`

	// key
	Key *string `json:"key,omitempty"`

	// 对象值
	ObjectValue *interface{} `json:"object_value,omitempty"`

	// 对象key
	ObjectKey *interface{} `json:"object_key,omitempty"`

	// 缺陷类型
	TrackerNames *interface{} `json:"tracker_names,omitempty"`

	// 归属看板信息，用例关联工作项信息使用
	BoardInfo *[]interface{} `json:"board_info,omitempty"`
}

func (o ElementResourceChangeExternalVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ElementResourceChangeExternalVo struct{}"
	}

	return strings.Join([]string{"ElementResourceChangeExternalVo", string(data)}, " ")
}
