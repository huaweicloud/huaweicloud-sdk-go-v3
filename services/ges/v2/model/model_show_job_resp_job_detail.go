package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 只有jobName为ImportGraph时才返回，用于显示导入图详情。
type ShowJobRespJobDetail struct {

	// 元数据路径。
	SchemaPath *[]ShowJobRespJobDetailSchemaPath `json:"schema_path,omitempty"`

	// 边数据集路径。
	EdgesetPath *[]ShowJobRespJobDetailSchemaPath `json:"edgeset_path,omitempty"`

	// 点数据集路径。
	VertexsetPath *[]ShowJobRespJobDetailSchemaPath `json:"vertexset_path,omitempty"`
}

func (o ShowJobRespJobDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobRespJobDetail struct{}"
	}

	return strings.Join([]string{"ShowJobRespJobDetail", string(data)}, " ")
}
