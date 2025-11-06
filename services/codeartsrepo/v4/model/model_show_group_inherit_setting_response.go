package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGroupInheritSettingResponse Response Object
type ShowGroupInheritSettingResponse struct {

	// 代码组继承设置
	Body           *[]ProjectSettingsInheritCfgDto `json:"body,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ShowGroupInheritSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupInheritSettingResponse struct{}"
	}

	return strings.Join([]string{"ShowGroupInheritSettingResponse", string(data)}, " ")
}
