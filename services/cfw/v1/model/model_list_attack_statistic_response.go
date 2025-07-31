package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAttackStatisticResponse Response Object
type ListAttackStatisticResponse struct {

	// **参数解释**： 攻击日志统计 **取值范围**： 不涉及
	Data           *[]AttackStatisticRespBody `json:"data,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListAttackStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAttackStatisticResponse struct{}"
	}

	return strings.Join([]string{"ListAttackStatisticResponse", string(data)}, " ")
}
