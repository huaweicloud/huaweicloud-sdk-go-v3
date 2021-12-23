package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 任务详情
type JobDetail struct {
	// 元数据路径。

	SchemaPath *[]SchemaPath `json:"schemaPath,omitempty"`
	// 边数据集路径。

	EdgesetPath *[]EdgesetPath `json:"edgesetPath,omitempty"`
	// 点数据集路径。

	VertexsetPath *[]VertexsetPath `json:"vertexsetPath,omitempty"`
}

func (o JobDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobDetail struct{}"
	}

	return strings.Join([]string{"JobDetail", string(data)}, " ")
}
