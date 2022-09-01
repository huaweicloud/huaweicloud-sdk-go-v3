package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateDisasterRecoveryDrillNameResponse struct {
	DisasterRecoveryDrill *ShowDisasterRecoveryDrillParams `json:"disaster_recovery_drill,omitempty" xml:"disaster_recovery_drill"`
	HttpStatusCode        int                              `json:"-"`
}

func (o UpdateDisasterRecoveryDrillNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDisasterRecoveryDrillNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateDisasterRecoveryDrillNameResponse", string(data)}, " ")
}
