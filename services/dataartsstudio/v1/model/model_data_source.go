package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataSource struct {

	// 数据连接名称
	DatasourceName *string `json:"datasource_name,omitempty"`

	// 数据连接类型
	DatasourceType *string `json:"datasource_type,omitempty"`

	// 数据连接guid
	DatasourceGuid *string `json:"datasource_guid,omitempty"`

	// 数据连接唯一标识名称
	DatasourceQualifiedName *string `json:"datasource_qualified_name,omitempty"`

	// obs目录数
	ObsFolderCount *int32 `json:"obs_folder_count,omitempty"`

	// obs文件数
	ObsFileCount *int32 `json:"obs_file_count,omitempty"`

	// css索引数
	CssIndexCount *int32 `json:"css_index_count,omitempty"`

	// css 索引字段数目
	CssIndexFieldCount *int32 `json:"css_index_field_count,omitempty"`

	// 命名空间数
	NamespaceCount *int32 `json:"namespace_count,omitempty"`

	// ges点的总数
	GesVertexCount *int32 `json:"ges_vertex_count,omitempty"`

	// ges边的总数
	GesEdgeCount *int32 `json:"ges_edge_count,omitempty"`

	// 数据库总数
	DatabaseCount *int32 `json:"database_count,omitempty"`

	// 通道总数
	StreamCount *int32 `json:"stream_count,omitempty"`

	// 表总数
	TableCount *int32 `json:"table_count,omitempty"`

	// 数据大小
	DataSize *float64 `json:"data_size,omitempty"`

	// 数据库统计信息
	Databases *[]Database `json:"databases,omitempty"`

	// 顶层目录统计信息
	Folders *[]ObsFolder `json:"folders,omitempty"`

	// css索引统计信息
	CssIndices *[]CssIndex `json:"css_indices,omitempty"`

	// 命名空间统计信息
	Namespaces *[]Namespace `json:"namespaces,omitempty"`

	// 通道统计信息
	DisStreams *[]DisStream `json:"dis_streams,omitempty"`
}

func (o DataSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataSource struct{}"
	}

	return strings.Join([]string{"DataSource", string(data)}, " ")
}
