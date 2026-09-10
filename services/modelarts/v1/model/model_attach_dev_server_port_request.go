package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachDevServerPortRequest Request Object
type AttachDevServerPortRequest struct {

	// **参数解释**：DevServer实例ID。 **约束限制**：必填。 **取值范围**：1 - 64字符。 **默认取值**：不涉及。
	Id string `json:"id"`

	Body *AttachDevServerPortsRequestBody `json:"body,omitempty"`
}

func (o AttachDevServerPortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachDevServerPortRequest struct{}"
	}

	return strings.Join([]string{"AttachDevServerPortRequest", string(data)}, " ")
}
