package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGroupNoteRequiredAttributesResponse Response Object
type UpdateGroupNoteRequiredAttributesResponse struct {

	// **参数解释：** 检视意见必填项。
	NoteRequiredAttributes *[]RequiredAttributeDto `json:"note_required_attributes,omitempty"`
	HttpStatusCode         int                     `json:"-"`
}

func (o UpdateGroupNoteRequiredAttributesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGroupNoteRequiredAttributesResponse struct{}"
	}

	return strings.Join([]string{"UpdateGroupNoteRequiredAttributesResponse", string(data)}, " ")
}
