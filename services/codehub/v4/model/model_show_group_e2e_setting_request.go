package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGroupE2eSettingRequest Request Object
type ShowGroupE2eSettingRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`
}

func (o ShowGroupE2eSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupE2eSettingRequest struct{}"
	}

	return strings.Join([]string{"ShowGroupE2eSettingRequest", string(data)}, " ")
}
