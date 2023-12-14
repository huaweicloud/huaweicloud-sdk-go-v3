package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OpenTsdbValue CloudTable集群OpenTSDB 数据value的Schema配置，用于将通道内的JSON数据进行格式转换生成OpenTSDB 数据的value。
type OpenTsdbValue struct {

	// 通道内用户JSON数据对应JSON属性的类型名称。  取值范围：  - Bigint - Double - Boolean - Timestamp - String - Decimal
	Type string `json:"type"`

	// 常量或通道内用户数据的JSON属性名称。  取值范围：1～32，只能包含英文字母、数字和下划线。
	Value string `json:"value"`
}

func (o OpenTsdbValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenTsdbValue struct{}"
	}

	return strings.Join([]string{"OpenTsdbValue", string(data)}, " ")
}
