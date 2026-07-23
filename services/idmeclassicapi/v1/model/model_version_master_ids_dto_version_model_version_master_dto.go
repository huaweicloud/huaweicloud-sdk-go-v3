package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionMasterIdsDtoVersionModelVersionMasterDto struct {

	// **参数解释：**  主对象集合，每个元素包含一个主对象ID及其对应的版本标识，用于批量定位待删除的分支。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	MasterIds []VersionModelMasterIdsDto `json:"masterIds"`

	// **参数解释：**  更新者账号，记录执行批量删除操作的用户标识。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`
}

func (o VersionMasterIdsDtoVersionModelVersionMasterDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionMasterIdsDtoVersionModelVersionMasterDto struct{}"
	}

	return strings.Join([]string{"VersionMasterIdsDtoVersionModelVersionMasterDto", string(data)}, " ")
}
