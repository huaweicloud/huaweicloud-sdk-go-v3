package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoPersistableModelCreateDto struct {
	Params *PersistableModelCreateDto `json:"params,omitempty"`

	// **参数解释：**  应用ID。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ApplicationId *string `json:"applicationId,omitempty"`
}

func (o RdmParamVoPersistableModelCreateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoPersistableModelCreateDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoPersistableModelCreateDto", string(data)}, " ")
}
