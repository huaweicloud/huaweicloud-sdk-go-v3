package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionMasterIdsDtoVersionModelVersionMasterDto struct {

	// **参数解释：**  主对象集合。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	MasterIds []VersionModelMasterIdsDto `json:"masterIds"`

	// **参数解释：**  更新者。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`
}

func (o VersionMasterIdsDtoVersionModelVersionMasterDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionMasterIdsDtoVersionModelVersionMasterDto struct{}"
	}

	return strings.Join([]string{"VersionMasterIdsDtoVersionModelVersionMasterDto", string(data)}, " ")
}
