package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeSrKernelVersionRequestV3 StarRocks内核版本升级请求体。
type UpgradeSrKernelVersionRequestV3 struct {

	// **参数解释**： 是否延时升级。  **约束限制**： 不涉及  **取值范围**： - true - false  **默认取值**： false。
	Delay *string `json:"delay,omitempty"`

	// **参数解释**： 是否跳过升级校验。  **约束限制**： 不涉及  **取值范围**： - true - false  **默认取值**： false。
	IsSkipValidate *string `json:"is_skip_validate,omitempty"`
}

func (o UpgradeSrKernelVersionRequestV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeSrKernelVersionRequestV3 struct{}"
	}

	return strings.Join([]string{"UpgradeSrKernelVersionRequestV3", string(data)}, " ")
}
