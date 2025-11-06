package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGroupMergeRequestApproverSettingRequest Request Object
type DeleteGroupMergeRequestApproverSettingRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`

	// **参数解释：** 合并请求审核设置id。
	SettingId int32 `json:"setting_id"`
}

func (o DeleteGroupMergeRequestApproverSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGroupMergeRequestApproverSettingRequest struct{}"
	}

	return strings.Join([]string{"DeleteGroupMergeRequestApproverSettingRequest", string(data)}, " ")
}
