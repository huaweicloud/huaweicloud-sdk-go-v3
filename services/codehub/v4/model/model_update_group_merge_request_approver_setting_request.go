package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGroupMergeRequestApproverSettingRequest Request Object
type UpdateGroupMergeRequestApproverSettingRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`

	// **参数解释：** 合并请求审核设置id。
	SettingId int32 `json:"setting_id"`

	Body *CreateMergeRequestApproverSettingDto `json:"body,omitempty"`
}

func (o UpdateGroupMergeRequestApproverSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGroupMergeRequestApproverSettingRequest struct{}"
	}

	return strings.Join([]string{"UpdateGroupMergeRequestApproverSettingRequest", string(data)}, " ")
}
