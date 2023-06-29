package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObjectTypeInfo 源库中的对象类型信息。
type ObjectTypeInfo struct {

	// 是否选择全部object类型。取值为true时，不包含USER。
	IsSelectAllObjectsType bool `json:"is_select_all_objects_type"`

	// 需要评估的object类型列表。is_select_all_objects_type为false时必填。
	ObjectsTypeList *[]string `json:"objects_type_list,omitempty"`
}

func (o ObjectTypeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectTypeInfo struct{}"
	}

	return strings.Join([]string{"ObjectTypeInfo", string(data)}, " ")
}
