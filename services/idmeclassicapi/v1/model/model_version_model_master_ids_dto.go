package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelMasterIdsDto struct {

	// **参数解释：**  父模型ID。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	MasterId string `json:"masterId"`

	// **参数解释：**  版本对象。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Version *string `json:"version,omitempty"`
}

func (o VersionModelMasterIdsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelMasterIdsDto struct{}"
	}

	return strings.Join([]string{"VersionModelMasterIdsDto", string(data)}, " ")
}
