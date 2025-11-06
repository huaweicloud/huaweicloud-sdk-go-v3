package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGroupPermissionInheritEnabledResponse Response Object
type ShowGroupPermissionInheritEnabledResponse struct {

	// **参数解释：** 使用项目权限配置开关是否开启。
	PermissionInheritEnabled *bool `json:"permission_inherit_enabled,omitempty"`
	HttpStatusCode           int   `json:"-"`
}

func (o ShowGroupPermissionInheritEnabledResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupPermissionInheritEnabledResponse struct{}"
	}

	return strings.Join([]string{"ShowGroupPermissionInheritEnabledResponse", string(data)}, " ")
}
