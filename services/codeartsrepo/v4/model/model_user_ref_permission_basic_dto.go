package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UserRefPermissionBasicDto struct {

	// **参数解释：** 是否有权限。 **取值范围：** true：有权限，false：没权限
	HasPermission *bool `json:"has_permission,omitempty"`

	// **参数解释：** 是否保护分支。 **取值范围：** true：是保护分支，false：非保护分支
	IsProtect *bool `json:"is_protect,omitempty"`
}

func (o UserRefPermissionBasicDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserRefPermissionBasicDto struct{}"
	}

	return strings.Join([]string{"UserRefPermissionBasicDto", string(data)}, " ")
}
