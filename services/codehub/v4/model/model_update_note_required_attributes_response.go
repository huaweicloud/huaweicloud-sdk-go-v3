package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNoteRequiredAttributesResponse Response Object
type UpdateNoteRequiredAttributesResponse struct {

	// **参数解释：** 检视意见必填项。
	NoteRequiredAttributes *[]RequiredAttributeDto `json:"note_required_attributes,omitempty"`
	HttpStatusCode         int                     `json:"-"`
}

func (o UpdateNoteRequiredAttributesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNoteRequiredAttributesResponse struct{}"
	}

	return strings.Join([]string{"UpdateNoteRequiredAttributesResponse", string(data)}, " ")
}
