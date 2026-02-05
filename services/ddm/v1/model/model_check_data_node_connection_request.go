package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckDataNodeConnectionRequest Request Object
type CheckDataNodeConnectionRequest struct {

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in09，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	Body *EsdbCheckRdsConnectionsRequestV3 `json:"body,omitempty"`
}

func (o CheckDataNodeConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDataNodeConnectionRequest struct{}"
	}

	return strings.Join([]string{"CheckDataNodeConnectionRequest", string(data)}, " ")
}
