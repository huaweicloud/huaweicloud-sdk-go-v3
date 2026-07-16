package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServiceSecret **参数解释：** 服务密钥挂载。 **约束限制：** 不涉及。
type ServiceSecret struct {

	// **参数解释：** 是否启用密钥。 **约束限制：** 不涉及。 **取值范围：** - true：启用密钥。 - false：不启用密钥。 **默认取值：** false。
	SecretEnable *bool `json:"secret_enable,omitempty"`

	// **参数解释：** 密钥类型。 **约束限制：** 不涉及。 **取值范围：** - custom：自定义密钥。 - [dew：DEW密钥。](tag:hws,hws_hk,fcs) **默认取值：** 不涉及。
	SecretType *string `json:"secret_type,omitempty"`

	// **参数解释：** 密钥挂载。 **约束限制：** 上限30个。
	SecretVolumes *[]SecretVolume `json:"secret_volumes,omitempty"`

	// **参数解释：** 是否启用镜像的用户组。 **约束限制：** 不涉及。 **取值范围：** - true：启用镜像的用户组。 - false：不启用镜像的用户组。 **默认取值：** false。
	GroupEnable *bool `json:"group_enable,omitempty"`

	// **参数解释：** 镜像的用户组ID。 **约束限制：** 不涉及。 **取值范围：** 1000~4294967294。 **默认取值：** 不涉及。
	GroupId *int64 `json:"group_id,omitempty"`
}

func (o ServiceSecret) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceSecret struct{}"
	}

	return strings.Join([]string{"ServiceSecret", string(data)}, " ")
}
