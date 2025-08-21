package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGroupWatermarkRequest Request Object
type UpdateGroupWatermarkRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`

	Body *UpdateWatermarkDto `json:"body,omitempty"`
}

func (o UpdateGroupWatermarkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGroupWatermarkRequest struct{}"
	}

	return strings.Join([]string{"UpdateGroupWatermarkRequest", string(data)}, " ")
}
