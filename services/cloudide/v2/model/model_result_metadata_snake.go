package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResultMetadataSnake struct {

	// 元数据类型
	MetadataType *string `json:"metadata_type,omitempty"`

	// 元数据列表
	MetadataItems *[]map[string]interface{} `json:"metadata_items,omitempty"`
}

func (o ResultMetadataSnake) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultMetadataSnake struct{}"
	}

	return strings.Join([]string{"ResultMetadataSnake", string(data)}, " ")
}
