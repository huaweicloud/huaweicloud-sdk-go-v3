package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportGraphReq 导出图请求体
type ExportGraphReq struct {

	// 图的导出OBS路径。
	GraphExportPath string `json:"graph_export_path"`

	// 导出边文件名。
	EdgeSetName string `json:"edge_set_name"`

	// 导出点文件名。
	VertexSetName string `json:"vertex_set_name"`

	// 导出元数据文件名。
	SchemaName string `json:"schema_name"`
}

func (o ExportGraphReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportGraphReq struct{}"
	}

	return strings.Join([]string{"ExportGraphReq", string(data)}, " ")
}
