package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpgradePathsResult struct {

	// **参数解释**: 源引擎版本号。 **取值范围**: 不涉及
	StartVersion *string `json:"start_version,omitempty"`

	// **参数解释**: 目标引擎版本号。 **取值范围**: 不涉及
	EndVersion *string `json:"end_version,omitempty"`
}

func (o UpgradePathsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradePathsResult struct{}"
	}

	return strings.Join([]string{"UpgradePathsResult", string(data)}, " ")
}
