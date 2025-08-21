package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGroupNoteRequiredAttributesRequest Request Object
type UpdateGroupNoteRequiredAttributesRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`

	Body *PostNoteRequiredAttributeDto `json:"body,omitempty"`
}

func (o UpdateGroupNoteRequiredAttributesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGroupNoteRequiredAttributesRequest struct{}"
	}

	return strings.Join([]string{"UpdateGroupNoteRequiredAttributesRequest", string(data)}, " ")
}
