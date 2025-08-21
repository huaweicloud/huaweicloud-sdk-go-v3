package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGroupNoteRequiredAttributesResponse Response Object
type ShowGroupNoteRequiredAttributesResponse struct {
	Body           *[]NoteRequiredAttributeDto `json:"body,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowGroupNoteRequiredAttributesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupNoteRequiredAttributesResponse struct{}"
	}

	return strings.Join([]string{"ShowGroupNoteRequiredAttributesResponse", string(data)}, " ")
}
