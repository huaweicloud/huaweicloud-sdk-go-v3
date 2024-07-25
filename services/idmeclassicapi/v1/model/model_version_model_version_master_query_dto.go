package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelVersionMasterQueryDto struct {

	// **参数解释：**  是否加密。  **取值范围：**  - true：加密。 - false：不加密。  **默认取值：**  false。
	Decrypt *bool `json:"decrypt,omitempty"`

	// **参数解释：**  迭代版本。如果此参数为空，则返回M-V模型实例的最新版本信息。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Iteration *int32 `json:"iteration,omitempty"`

	// **参数解释：**  主对象ID。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	MasterId string `json:"masterId"`

	// **参数解释：**  版本号。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Version string `json:"version"`
}

func (o VersionModelVersionMasterQueryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelVersionMasterQueryDto struct{}"
	}

	return strings.Join([]string{"VersionModelVersionMasterQueryDto", string(data)}, " ")
}
