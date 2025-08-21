package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectMemberSettingResponse Response Object
type ShowProjectMemberSettingResponse struct {

	// **参数解释：** 项目id **取值范围：** 字符串长度不少于1，不超过1000。
	ProductId *string `json:"product_id,omitempty"`

	// **参数解释：** 同步开关
	SyncEnabled *bool `json:"sync_enabled,omitempty"`

	// **参数解释：** 同步所有角色开关
	SyncAllRoleEnabled *bool `json:"sync_all_role_enabled,omitempty"`

	// **参数解释：** 角色同步
	RoleSync       *[]RoleSyncDto `json:"role_sync,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowProjectMemberSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectMemberSettingResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectMemberSettingResponse", string(data)}, " ")
}
