package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNoteRequestBody 创建评论请求体
type CreateNoteRequestBody struct {

	// 评论的对象ID
	WarRoomId *string `json:"war_room_id,omitempty"`

	// 评论的类型
	Type string `json:"type"`

	// 评论的内容
	Content string `json:"content"`

	// 评论的动作类型
	NoteType *string `json:"note_type,omitempty"`
}

func (o CreateNoteRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNoteRequestBody struct{}"
	}

	return strings.Join([]string{"CreateNoteRequestBody", string(data)}, " ")
}
