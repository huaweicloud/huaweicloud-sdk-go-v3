package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 增量导入图请求体
type ImportGraphReq struct {

	// 边文件目录或边文件名。
	EdgesetPath *string `json:"edgeset_path,omitempty"`

	// 边数据集格式。当前仅支持csv。  默认为csv。
	EdgesetFormat *string `json:"edgeset_format,omitempty"`

	// 点文件目录或点文件名。
	VertexsetPath *string `json:"vertexset_path,omitempty"`

	// 点数据集格式。当前仅支持csv。  默认为csv。
	VertexsetFormat *string `json:"vertexset_format,omitempty"`

	// 新增数据的元数据文件路径。
	SchemaPath *string `json:"schema_path,omitempty"`

	// 导入图日志存放目录，用于存储导入失败的数据和详细错入原因。
	LogDir *string `json:"log_dir,omitempty"`

	ParallelEdge *ImportGraphReqParallelEdge `json:"parallel_edge,omitempty"`

	// csv格式文件字段分隔符，默认值为逗号（,）。list/set类型的字段内元素分隔符默认为分号（;）。
	Delimiter *string `json:"delimiter,omitempty"`

	// csv格式文件字段包围符，默认值为双引号（\"）。用来包围一个字段，如字段中含有分隔符或者换行等。
	TrimQuote *string `json:"trim_quote,omitempty"`

	// 是否离线导入，取值为true或者false，默认取false。  - true 表示离线导入，导入速度较快，但导入过程中图处于锁定状态，不可读不可写。 - false 表示在线导入，相对离线导入，在线导入速度略慢，但导入过程中图并未锁定，可读不可写。
	Offline *bool `json:"offline,omitempty"`
}

func (o ImportGraphReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportGraphReq struct{}"
	}

	return strings.Join([]string{"ImportGraphReq", string(data)}, " ")
}
