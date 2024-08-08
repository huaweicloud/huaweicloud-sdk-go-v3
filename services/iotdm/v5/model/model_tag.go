package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Tag struct {

	// **参数说明**：标签键，标签的键可用字母(包含中文)、数字、空格和以下字符：下划线(_)点(.)冒号(:)等号(=)加号(+)中划线(-)以及@，首尾不能有空格字符
	Key string `json:"key"`

	// **参数说明**：标签值，可为空字符串和null，标签的值可用字母(包含中文)、数字、空格和以下字符：下划线(_)点(.)冒号(:)等号(=)加号(+)中划线(-)以及@
	Value *string `json:"value,omitempty"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
