package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDisasterInfoRequest Request Object
type UpdateDisasterInfoRequest struct {

	// **参数解释**： 容灾ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DisasterRecoveryId string `json:"disaster_recovery_id"`

	Body *UpdateDisasterRecoveryRequest `json:"body,omitempty"`
}

func (o UpdateDisasterInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDisasterInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateDisasterInfoRequest", string(data)}, " ")
}
