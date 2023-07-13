package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportGraphReq This is a auto create Body Object
type ImportGraphReq struct {

	// 边文件目录或边文件名。
	EdgesetPath *string `json:"edgesetPath,omitempty"`

	// 边数据集格式。当前仅支持csv。  默认为csv。
	EdgesetFormat *string `json:"edgesetFormat,omitempty"`

	// 点文件目录或点文件名。
	VertexsetPath *string `json:"vertexsetPath,omitempty"`

	// 点数据集格式。当前仅支持csv。  默认为csv。
	VertexsetFormat *string `json:"vertexsetFormat,omitempty"`

	// 新增数据的元数据文件路径。
	SchemaPath *string `json:"schemaPath,omitempty"`

	// 导入图日志存放目录，用于存储导入失败的数据和详细错入原因。
	LogDir *string `json:"logDir,omitempty"`

	// 重复边处理
	ParallelEdge *interface{} `json:"parallelEdge,omitempty"`

	// 处理方式，取值为allow，ignore和override，默认为allow。 - allow表示允许重复边。 - ignore表示忽略之后的重复边。 - override表示覆盖之前的重复边。
	Action *string `json:"action,omitempty"`

	// 重复边的定义，是否忽略Label。取值为true或者false，默认取true。 - true 表示重复边定义不包含Label，即用<源点，终点>标记一条边，不包含Label。 - false 表示重复边定义包含Label，即用<源点，终点，Label>标记一条边。
	IgnoreLabel *bool `json:"ignoreLabel,omitempty"`

	// csv格式文件字段分隔符，默认值为逗号（,）。list/set类型的字段内元素分隔符默认为分号（;）。
	Delimiter *string `json:"delimiter,omitempty"`

	// csv格式文件字段包围符，默认值为双引号（\"）。用来包围一个字段，如字段中含有分隔符或者换行等。
	TrimQuote *string `json:"trimQuote,omitempty"`

	// 是否离线导入，取值为true或者false，默认取false。 - true 表示离线导入，导入速度较快，但导入过程中图处于锁定状态，不可读不可写。 - false 表示在线导入，相对离线导入，在线导入速度略慢，但导入过程中图并未锁定，可读不可写。
	Offline *bool `json:"offline,omitempty"`
}

func (o ImportGraphReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportGraphReq struct{}"
	}

	return strings.Join([]string{"ImportGraphReq", string(data)}, " ")
}
