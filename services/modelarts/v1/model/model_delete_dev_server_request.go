package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDevServerRequest Request Object
type DeleteDevServerRequest struct {

	// **参数解释**：Lite Server ID。 **约束限制**：必填。 **取值范围**：1 - 64字符。 **默认取值**：不涉及。
	Id string `json:"id"`
}

func (o DeleteDevServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDevServerRequest struct{}"
	}

	return strings.Join([]string{"DeleteDevServerRequest", string(data)}, " ")
}
