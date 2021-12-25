package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标签属性
type Tag struct {
	// 标签key值，长度取值范围为1~36， 仅允许大小写英文字母、数字、下划线、中划线

	Key string `json:"key"`
	// 标签value值，长度取值范围为0~43， 仅允许大小写英文字母、数字、下划线、中划线

	Values *[]string `json:"values,omitempty"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
