package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OpenTsdbTimestamp CloudTable集群OpenTSDB 数据timestamp的Schema配置，用于将通道内的JSON数据进行格式转换生成OpenTSDB数据的timestamp。
type OpenTsdbTimestamp struct {

	// - Timestamp类型表示通道内用户数据对应JSON属性的取值为Timestamp类型，不需要进行数据格式转换就可以生成OpenTSDB的timestamp。 - String类型表示通道内用户数据对应JSON属性的取值为Date格式，需要进行数据格式转换才能生成OpenTSDB的timestamp。
	Type string `json:"type"`

	// 通道内用户数据的JSON属性名称。  取值范围：1～32，只能包含英文字母、数字和下划线。
	Value string `json:"value"`

	// “type”为“String”类型时必选。表示通道内用户数据对应JSON属性的取值为Date格式，需要根据format字段进行数据格式转换生成OpenTSDB的timestamp。  取值范围：  - yyyy/MM/dd HH:mm:ss - MM/dd/yyyy HH:mm:ss - dd/MM/yyyy HH:mm:ss - yyyy-MM-dd HH:mm:ss - MM-dd-yyyy HH:mm:ss - dd-MM-yyyy HH:mm:ss
	Format string `json:"format"`
}

func (o OpenTsdbTimestamp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenTsdbTimestamp struct{}"
	}

	return strings.Join([]string{"OpenTsdbTimestamp", string(data)}, " ")
}
