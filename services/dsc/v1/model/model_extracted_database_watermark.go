package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtractedDatabaseWatermark struct {

	// 水印密钥
	WatermarkKey string `json:"watermark_key"`

	// 字段类型列表，最大长度100。使用时，要包含嵌入时所有“primary_key”为true的字段，和至少一个为false的字段用来提取水印
	Columns []Columns `json:"columns"`

	// 水印数据，数据条数不超过30,000条
	Data []map[string]interface{} `json:"data"`
}

func (o ExtractedDatabaseWatermark) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtractedDatabaseWatermark struct{}"
	}

	return strings.Join([]string{"ExtractedDatabaseWatermark", string(data)}, " ")
}
