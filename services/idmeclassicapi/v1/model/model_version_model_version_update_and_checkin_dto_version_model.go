package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelVersionUpdateAndCheckinDtoVersionModel struct {
	Data *VersionModel `json:"data"`

	// **参数解释：**  主对象ID，用于标识待更新并检入的M-V模型实例的主对象。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	MasterId string `json:"masterId"`

	// **参数解释：**  更新者账号。  **约束限制：**  需与实例当前modifier字段值一致，否则将触发权限校验失败。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`
}

func (o VersionModelVersionUpdateAndCheckinDtoVersionModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelVersionUpdateAndCheckinDtoVersionModel struct{}"
	}

	return strings.Join([]string{"VersionModelVersionUpdateAndCheckinDtoVersionModel", string(data)}, " ")
}
