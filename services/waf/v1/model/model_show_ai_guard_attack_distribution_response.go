package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAiGuardAttackDistributionResponse Response Object
type ShowAiGuardAttackDistributionResponse struct {

	// 总命中数，数量等于下一级items的total的数量
	Num *int64 `json:"num,omitempty"`

	// 攻击类型分布列表
	Items *[]AttackDistributionItem `json:"items,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAiGuardAttackDistributionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAiGuardAttackDistributionResponse struct{}"
	}

	return strings.Join([]string{"ShowAiGuardAttackDistributionResponse", string(data)}, " ")
}
