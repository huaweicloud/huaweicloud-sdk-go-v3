package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpenTsdbTags struct {

	// 存储该通道数据的OpenTSDB数据的tag名称。  取值范围：1～32，只能包含英文字母、数字和下划线。
	Name string `json:"name"`

	// 通道内用户JSON数据对应JSON属性的类型名称。  取值范围：  - Bigint - Double - Boolean - Timestamp - String - Decimal
	Type string `json:"type"`

	// 常量或通道内用户数据的JSON属性名称。  取值范围：1～32，只能包含英文字母、数字和下划线。
	Value string `json:"value"`
}

func (o OpenTsdbTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenTsdbTags struct{}"
	}

	return strings.Join([]string{"OpenTsdbTags", string(data)}, " ")
}
