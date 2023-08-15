package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtensionQueryResult struct {

	// 插件列表集合
	Extensions *[]ExtensionAllSnake `json:"extensions,omitempty"`

	// 结果元数据集合
	ResultMetadata *[]ResultMetadataSnake `json:"result_metadata,omitempty"`
}

func (o ExtensionQueryResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionQueryResult struct{}"
	}

	return strings.Join([]string{"ExtensionQueryResult", string(data)}, " ")
}
