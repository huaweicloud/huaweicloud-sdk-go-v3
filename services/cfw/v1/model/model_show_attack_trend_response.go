package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAttackTrendResponse Response Object
type ShowAttackTrendResponse struct {

	// **参数解释**： 攻击趋势 **取值范围**： 不涉及
	Data           *[]AttackTrendRespBody `json:"data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowAttackTrendResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAttackTrendResponse struct{}"
	}

	return strings.Join([]string{"ShowAttackTrendResponse", string(data)}, " ")
}
