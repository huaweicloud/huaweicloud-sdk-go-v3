package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PositionDto struct {

	// 源分支base提交节点
	BaseSha *string `json:"base_sha,omitempty"`

	// 目标分支最新提交节点
	StartSha *string `json:"start_sha,omitempty"`

	// 源分支最新提交节点
	HeadSha *string `json:"head_sha,omitempty"`

	// 修改前文件路径
	OldPath *string `json:"old_path,omitempty"`

	// 修改后文件路径
	NewPath *string `json:"new_path,omitempty"`

	// 变更类型
	PositionType *string `json:"position_type,omitempty"`

	// 修改前行号
	OldLine *int32 `json:"old_line,omitempty"`

	// 修改后行号
	NewLine *int32 `json:"new_line,omitempty"`
}

func (o PositionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PositionDto struct{}"
	}

	return strings.Join([]string{"PositionDto", string(data)}, " ")
}
