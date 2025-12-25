package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompatibleDataStoreResp **参数解释**： 版本信息。 **取值范围**： 不涉及。
type CompatibleDataStoreResp struct {

	// **参数解释**： 数据库类型。 **取值范围**： - dws：云数仓 - hybrid：实时数仓
	Type *string `json:"type,omitempty"`

	// **参数解释**： 版本。 **取值范围**： 不涉及。
	Version *string `json:"version,omitempty"`
}

func (o CompatibleDataStoreResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompatibleDataStoreResp struct{}"
	}

	return strings.Join([]string{"CompatibleDataStoreResp", string(data)}, " ")
}
