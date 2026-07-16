package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecretVolume **参数解释：** 密钥挂载。 **约束限制：** 不涉及。
type SecretVolume struct {

	// **参数解释：** 密钥名称。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	SecretName *string `json:"secret_name,omitempty"`

	// **参数解释：** 密钥key。 **约束限制：** 匹配一个长度在1到63之间的字符串，只能包含字母、数字、点、下划线和连字符，并且不能以两个连续的点（..）开头。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	SecretKey *string `json:"secret_key,omitempty"`

	// **参数解释：** 密钥值。 **约束限制：** 长度在1~32768，Base64编码后的密钥值。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	SecretValue *string `json:"secret_value,omitempty"`

	// **参数解释：** 挂载路径。 **约束限制：** 不涉及。 **取值范围：** 以(/)开头和结尾，可包含字母、数字、中划线、下划线，整个挂载路径长度不能超过255位。 **默认取值：** 不涉及。
	MountPath string `json:"mount_path"`
}

func (o SecretVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecretVolume struct{}"
	}

	return strings.Join([]string{"SecretVolume", string(data)}, " ")
}
