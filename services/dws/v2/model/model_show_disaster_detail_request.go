package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDisasterDetailRequest Request Object
type ShowDisasterDetailRequest struct {

	// **参数解释**： 容灾ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DisasterRecoveryId string `json:"disaster_recovery_id"`
}

func (o ShowDisasterDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDisasterDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowDisasterDetailRequest", string(data)}, " ")
}
