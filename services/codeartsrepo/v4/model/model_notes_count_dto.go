package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NotesCountDto struct {

	// 检视意见总数
	NotesCount *int32 `json:"notes_count,omitempty"`

	// 未解决的检视意见数量
	UnresolvedNotesCount *int32 `json:"unresolved_notes_count,omitempty"`

	// 已解决的检视意见数量
	AlreadyResolvedCount *int32 `json:"already_resolved_count,omitempty"`

	// 需要解决的检视意见总数
	NeedResolvedCount *int32 `json:"need_resolved_count,omitempty"`
}

func (o NotesCountDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotesCountDto struct{}"
	}

	return strings.Join([]string{"NotesCountDto", string(data)}, " ")
}
