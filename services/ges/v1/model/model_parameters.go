package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Parameters struct {

	// 元数据文件OBS路径，只支持文件。
	SchemaPath string `json:"schemaPath" xml:"schemaPath"`

	// 边数据集文件OBS路径，只支持文件。
	EdgesetPath string `json:"edgesetPath" xml:"edgesetPath"`

	// 边数据集格式。当前仅支持csv。  默认为csv。
	EdgesetFormat *string `json:"edgesetFormat,omitempty" xml:"edgesetFormat"`

	// 边数据集默认标签，当前默认为空，可以不填。
	EdgesetDefaultLabel *string `json:"edgesetDefaultLabel,omitempty" xml:"edgesetDefaultLabel"`

	// 点数据集OBS路径，只支持文件。
	VertexsetPath *string `json:"vertexsetPath,omitempty" xml:"vertexsetPath"`

	// 点数据集格式。当前仅支持csv。  默认为csv。
	VertexsetFormat *string `json:"vertexsetFormat,omitempty" xml:"vertexsetFormat"`

	// 点数据集默认标签，当前默认为空，可以不填。
	VertexsetDefaultLabel *string `json:"vertexsetDefaultLabel,omitempty" xml:"vertexsetDefaultLabel"`

	// OBS日志存储目录，用于存储建图过程导入失败的数据和详细日志。
	LogDir *string `json:"logDir,omitempty" xml:"logDir"`

	ParallelEdge *ParallelEdge `json:"parallelEdge,omitempty" xml:"parallelEdge"`
}

func (o Parameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Parameters struct{}"
	}

	return strings.Join([]string{"Parameters", string(data)}, " ")
}
