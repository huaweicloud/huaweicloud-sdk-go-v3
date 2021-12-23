package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type CheckMetadataReq struct {
	// 元数据文件的OBS路径。

	SchemaPath string `json:"schemaPath"`
	// 边数据集的OBS路径。

	EdgesetPath string `json:"edgesetPath"`
	// 边数据集格式。当前仅支持csv。  默认为csv。

	EdgesetFormat *string `json:"edgesetFormat,omitempty"`
	// 点数据集格式。当前仅支持csv。  默认为csv。

	VertexsetFormat *string `json:"vertexsetFormat,omitempty"`
	// 点数据集OBS路径。

	VertexsetPath *string `json:"vertexsetPath,omitempty"`
}

func (o CheckMetadataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckMetadataReq struct{}"
	}

	return strings.Join([]string{"CheckMetadataReq", string(data)}, " ")
}
