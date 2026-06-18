package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SchemaDto **参数解释：** 代码导航功模式信息
type SchemaDto struct {

	// **参数解释：** 代码导航版本。 **约束限制：** 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释：** 支持的最大文件大小。 **约束限制：** 不涉及。
	MaximumFileSize *int32 `json:"maximum_file_size,omitempty"`

	// **参数解释：** 支持的最大行数。 **约束限制：** 不涉及。
	MaximumLineLength *int32 `json:"maximum_line_length,omitempty"`

	// **参数解释：** 每行支持的最大字符数，超过将截断。 **约束限制：** 不涉及。
	MaximumTruncateLine *int32 `json:"maximum_truncate_line,omitempty"`

	// **参数解释：** 索引创建时间。 **约束限制：** 不涉及。
	CreateAt *string `json:"create_at,omitempty"`

	// **参数解释：** 索引更新时间。 **约束限制：** 不涉及。
	UpdateAt *string `json:"update_at,omitempty"`

	// **参数解释：** 索引重建时间。 **约束限制：** 不涉及。
	RebuildAt *string `json:"rebuild_at,omitempty"`

	// **参数解释：** 索引最近构建时间。 **约束限制：** 不涉及。
	LastBuildAt *string `json:"last_build_at,omitempty"`

	// **参数解释：** 索引构建次数。 **约束限制：** 不涉及。
	BuildTimes *int32 `json:"build_times,omitempty"`

	// **参数解释：** 请求次数。 **约束限制：** 不涉及。
	QueryTimes *int32 `json:"query_times,omitempty"`

	// **参数解释：** 索引大纲请求次数。 **约束限制：** 不涉及。
	OutlineTimes *int32 `json:"outline_times,omitempty"`
}

func (o SchemaDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SchemaDto struct{}"
	}

	return strings.Join([]string{"SchemaDto", string(data)}, " ")
}
