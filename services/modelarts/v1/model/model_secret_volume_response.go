package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecretVolumeResponse **参数解释：** 密钥挂载。
type SecretVolumeResponse struct {

	// **参数解释：** 密钥名称。 **取值范围：** 不涉及。
	SecretName *string `json:"secret_name,omitempty"`

	// **参数解释：** 密钥key。 **取值范围：** 长度不大于63。
	SecretKey *string `json:"secret_key,omitempty"`

	// **参数解释：** 密钥值。 **取值范围：** 长度不大于32768。
	SecretValue *string `json:"secret_value,omitempty"`

	// **参数解释：** 挂载路径。 **取值范围：** 不涉及。
	MountPath string `json:"mount_path"`
}

func (o SecretVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecretVolumeResponse struct{}"
	}

	return strings.Join([]string{"SecretVolumeResponse", string(data)}, " ")
}
