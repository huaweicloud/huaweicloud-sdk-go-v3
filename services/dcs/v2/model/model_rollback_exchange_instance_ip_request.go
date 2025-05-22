package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RollbackExchangeInstanceIpRequest Request Object
type RollbackExchangeInstanceIpRequest struct {

	// **参数解释**： 数据迁移任务ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TaskId string `json:"task_id"`
}

func (o RollbackExchangeInstanceIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollbackExchangeInstanceIpRequest struct{}"
	}

	return strings.Join([]string{"RollbackExchangeInstanceIpRequest", string(data)}, " ")
}
