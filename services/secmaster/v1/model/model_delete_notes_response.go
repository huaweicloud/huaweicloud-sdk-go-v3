package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNotesResponse Response Object
type DeleteNotesResponse struct {

	// **参数解释**: 错误码 **取值范围**: 不涉及
	Code *string `json:"code,omitempty"`

	// **参数解释**: 错误描述 **取值范围**: 不涉及
	Message *string `json:"message,omitempty"`

	Data           *DeleteNoteResult `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteNotesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNotesResponse struct{}"
	}

	return strings.Join([]string{"DeleteNotesResponse", string(data)}, " ")
}
