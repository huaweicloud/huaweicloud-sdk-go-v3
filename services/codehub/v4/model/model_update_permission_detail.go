package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePermissionDetail struct {

	// **参数解释：** 权限点id **取值范围：** 不涉及。
	PermissionId *int32 `json:"permission_id,omitempty"`

	// **参数解释：** 权限点状态 **取值范围：** - true, 开启。 - false, 关闭。
	Enabled *bool `json:"enabled,omitempty"`
}

func (o UpdatePermissionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePermissionDetail struct{}"
	}

	return strings.Join([]string{"UpdatePermissionDetail", string(data)}, " ")
}
