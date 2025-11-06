package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGroupReviewSettingsRequest Request Object
type UpdateGroupReviewSettingsRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`

	Body *ReviewSettingParamDto `json:"body,omitempty"`
}

func (o UpdateGroupReviewSettingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGroupReviewSettingsRequest struct{}"
	}

	return strings.Join([]string{"UpdateGroupReviewSettingsRequest", string(data)}, " ")
}
