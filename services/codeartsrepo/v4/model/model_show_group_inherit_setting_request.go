package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGroupInheritSettingRequest Request Object
type ShowGroupInheritSettingRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`
}

func (o ShowGroupInheritSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupInheritSettingRequest struct{}"
	}

	return strings.Join([]string{"ShowGroupInheritSettingRequest", string(data)}, " ")
}
