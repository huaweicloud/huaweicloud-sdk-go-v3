package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelMasterIdsDto struct {

	// **参数解释：**  主对象ID，用于定位待删除分支所属的主对象。  **约束限制：**  单次请求不超过1000个。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	MasterId string `json:"masterId"`

	// **参数解释：**  版本对象，用于指定待删除的分支版本标识。若未指定，则默认删除该主对象下的最新分支。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Version *string `json:"version,omitempty"`
}

func (o VersionModelMasterIdsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelMasterIdsDto struct{}"
	}

	return strings.Join([]string{"VersionModelMasterIdsDto", string(data)}, " ")
}
