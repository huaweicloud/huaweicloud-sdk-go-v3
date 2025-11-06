package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PositionDto 检视意见位置详情。
type PositionDto struct {

	// **参数解释：** 基础sha值。
	BaseSha *string `json:"base_sha,omitempty"`

	// **参数解释：** 起始sha值。
	StartSha *string `json:"start_sha,omitempty"`

	// **参数解释：** 最新sha值。
	HeadSha *string `json:"head_sha,omitempty"`

	// **参数解释：** 文件旧路径。
	OldPath *string `json:"old_path,omitempty"`

	// **参数解释：** 文件新路径。
	NewPath *string `json:"new_path,omitempty"`

	// **参数解释：** 文件类型。
	PositionType *string `json:"position_type,omitempty"`

	// **参数解释：** 旧文件行号。
	OldLine *int32 `json:"old_line,omitempty"`

	// **参数解释：** 新文件行号。
	NewLine *int32 `json:"new_line,omitempty"`
}

func (o PositionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PositionDto struct{}"
	}

	return strings.Join([]string{"PositionDto", string(data)}, " ")
}
