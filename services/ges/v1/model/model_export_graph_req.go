package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type ExportGraphReq struct {

	// 图的导出OBS路径。
	GraphExportPath string `json:"graphExportPath"`

	// 导出边文件名。
	EdgeSetName string `json:"edgeSetName"`

	// 导出点文件名。
	VertexSetName string `json:"vertexSetName"`

	// 导出元数据文件名。
	SchemaName string `json:"schemaName"`
}

func (o ExportGraphReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportGraphReq struct{}"
	}

	return strings.Join([]string{"ExportGraphReq", string(data)}, " ")
}
