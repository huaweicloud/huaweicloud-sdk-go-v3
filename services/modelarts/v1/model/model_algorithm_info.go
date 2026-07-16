package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlgorithmInfo 算法详情
type AlgorithmInfo struct {

	// **参数解释**：算法id。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Id *string `json:"id,omitempty"`
}

func (o AlgorithmInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlgorithmInfo struct{}"
	}

	return strings.Join([]string{"AlgorithmInfo", string(data)}, " ")
}
