package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServiceSecretResponse **参数解释：** 服务密钥挂载。
type ServiceSecretResponse struct {

	// **参数解释：** 是否启用密钥。 **取值范围：** - true：启用密钥。 - false：不启用密钥。
	SecretEnable *bool `json:"secret_enable,omitempty"`

	// **参数解释：** 密钥类型。 **取值范围：** - custom：自定义密钥。 - [dew：DEW密钥。](tag:hws,hws_hk,fcs)
	SecretType *string `json:"secret_type,omitempty"`

	// **参数解释：** 密钥挂载。 **约束限制：** 上限30个。
	SecretVolumes *[]SecretVolumeResponse `json:"secret_volumes,omitempty"`

	// **参数解释：** 是否启用镜像的用户组。 **取值范围：** - true：启用镜像的用户组。 - false：不启用镜像的用户组。
	GroupEnable *bool `json:"group_enable,omitempty"`

	// **参数解释：** 镜像的用户组ID。 **取值范围：** 1000~4294967294。
	GroupId *int64 `json:"group_id,omitempty"`
}

func (o ServiceSecretResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceSecretResponse struct{}"
	}

	return strings.Join([]string{"ServiceSecretResponse", string(data)}, " ")
}
