package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNoteRequiredAttributesResponse Response Object
type ShowNoteRequiredAttributesResponse struct {
	Body           *[]NoteRequiredAttributeDto `json:"body,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowNoteRequiredAttributesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNoteRequiredAttributesResponse struct{}"
	}

	return strings.Join([]string{"ShowNoteRequiredAttributesResponse", string(data)}, " ")
}
