package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGroupMergeRequestApproverSettingRequest Request Object
type CreateGroupMergeRequestApproverSettingRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`

	Body *CreateMergeRequestApproverSettingDto `json:"body,omitempty"`
}

func (o CreateGroupMergeRequestApproverSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupMergeRequestApproverSettingRequest struct{}"
	}

	return strings.Join([]string{"CreateGroupMergeRequestApproverSettingRequest", string(data)}, " ")
}
