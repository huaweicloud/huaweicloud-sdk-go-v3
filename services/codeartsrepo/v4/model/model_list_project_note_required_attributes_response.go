package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectNoteRequiredAttributesResponse Response Object
type ListProjectNoteRequiredAttributesResponse struct {

	// **参数解释：** 检视意见必填项。
	NoteRequiredAttributes *[]RequiredAttributeDto `json:"note_required_attributes,omitempty"`
	HttpStatusCode         int                     `json:"-"`
}

func (o ListProjectNoteRequiredAttributesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectNoteRequiredAttributesResponse struct{}"
	}

	return strings.Join([]string{"ListProjectNoteRequiredAttributesResponse", string(data)}, " ")
}
