package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotesDetail 评论详情对象
type NotesDetail struct {

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	Data *NotesDetailData `json:"data,omitempty"`

	// 评论唯一UUID
	Id *string `json:"id,omitempty"`

	// 是否已删除
	IsDeleted *bool `json:"is_deleted,omitempty"`

	// 标识是否是评论
	MarkedNote *bool `json:"marked_note,omitempty"`

	// 评论的动作类型
	NoteType *string `json:"note_type,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 评论的类型
	Type *string `json:"type,omitempty"`

	User *NotesDetailUser `json:"user,omitempty"`

	// 评论的对象ID
	WarRoomId *string `json:"war_room_id,omitempty"`

	// 空间
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 评论详情信息
	Content *interface{} `json:"content,omitempty"`
}

func (o NotesDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotesDetail struct{}"
	}

	return strings.Join([]string{"NotesDetail", string(data)}, " ")
}
