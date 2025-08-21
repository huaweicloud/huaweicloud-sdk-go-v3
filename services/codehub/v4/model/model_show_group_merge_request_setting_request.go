package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGroupMergeRequestSettingRequest Request Object
type ShowGroupMergeRequestSettingRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`
}

func (o ShowGroupMergeRequestSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupMergeRequestSettingRequest struct{}"
	}

	return strings.Join([]string{"ShowGroupMergeRequestSettingRequest", string(data)}, " ")
}
