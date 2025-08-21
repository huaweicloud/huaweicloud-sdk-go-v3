package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGroupNoteRequiredAttributesRequest Request Object
type ShowGroupNoteRequiredAttributesRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`
}

func (o ShowGroupNoteRequiredAttributesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupNoteRequiredAttributesRequest struct{}"
	}

	return strings.Join([]string{"ShowGroupNoteRequiredAttributesRequest", string(data)}, " ")
}
