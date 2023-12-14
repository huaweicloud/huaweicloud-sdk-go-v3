package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CsvProperties CSV 格式数据的相关属性，比如分隔符 delimiter
type CsvProperties struct {

	// 数据分隔符。
	Delimiter *string `json:"delimiter,omitempty"`
}

func (o CsvProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CsvProperties struct{}"
	}

	return strings.Join([]string{"CsvProperties", string(data)}, " ")
}
