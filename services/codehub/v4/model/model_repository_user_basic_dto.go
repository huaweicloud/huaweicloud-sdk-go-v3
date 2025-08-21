package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryUserBasicDto struct {

	// **参数解释：**  用户id **约束限制：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：**  用户名称 **约束限制：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：**  用户名 **约束限制：** 不涉及。
	Username *string `json:"username,omitempty"`

	// **参数解释：**  用户状态 **约束限制：** 不涉及。
	State *string `json:"state,omitempty"`

	// **参数解释：**  服务级权限状态 0：停用 1：启用 **约束限制：** 不涉及。
	ServiceLicenseStatus *int32 `json:"service_license_status,omitempty"`

	// **参数解释：**  用户中文名称 **约束限制：** 不涉及。
	NameCn *string `json:"name_cn,omitempty"`

	// **参数解释：**  用户昵称 **约束限制：** 不涉及。
	NickName *string `json:"nick_name,omitempty"`

	// **参数解释：**  租户名称 **约束限制：** 不涉及。
	TenantName *string `json:"tenant_name,omitempty"`
}

func (o RepositoryUserBasicDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryUserBasicDto struct{}"
	}

	return strings.Join([]string{"RepositoryUserBasicDto", string(data)}, " ")
}
