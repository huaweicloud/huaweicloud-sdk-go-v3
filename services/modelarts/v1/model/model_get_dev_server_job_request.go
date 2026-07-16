package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetDevServerJobRequest Request Object
type GetDevServerJobRequest struct {

	// **参数解释**：Lite Server job id。 **约束限制**：必填。 **取值范围**：1 - 64字符。 **默认取值**：不涉及。
	Id string `json:"id"`
}

func (o GetDevServerJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetDevServerJobRequest struct{}"
	}

	return strings.Join([]string{"GetDevServerJobRequest", string(data)}, " ")
}
