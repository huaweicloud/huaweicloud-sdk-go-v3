package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNoteResponse Response Object
type CreateNoteResponse struct {

	// **参数解释**: 错误码 **取值范围**: 不涉及
	Code *string `json:"code,omitempty"`

	// **参数解释**: 错误描述 **取值范围**: 不涉及
	Message *string `json:"message,omitempty"`

	// 是否请求成功
	Success *bool `json:"success,omitempty"`

	Data           *NotesDetail `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CreateNoteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNoteResponse struct{}"
	}

	return strings.Join([]string{"CreateNoteResponse", string(data)}, " ")
}
