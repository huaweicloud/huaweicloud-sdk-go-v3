package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckRdsConnectionResqVo struct {

	// **参数解释**：  rds实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in01，长度为36个字符。  **默认取值**：  不涉及。
	RdsInstanceId *string `json:"rds_instance_id,omitempty"`

	// success
	Success *string `json:"success,omitempty"`

	// **参数解释**：  错误码。  **参数范围**：  不涉及。
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释**：  错误信息。  **参数范围**：  不涉及。
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o CheckRdsConnectionResqVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRdsConnectionResqVo struct{}"
	}

	return strings.Join([]string{"CheckRdsConnectionResqVo", string(data)}, " ")
}
