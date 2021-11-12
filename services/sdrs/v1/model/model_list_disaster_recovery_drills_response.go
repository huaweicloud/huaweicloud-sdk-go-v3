package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDisasterRecoveryDrillsResponse struct {
	// 容灾演练列表。

	DisasterRecoveryDrills *[]ShowDisasterRecoveryDrillParams `json:"disaster_recovery_drills,omitempty"`
	// 列表中包含的容灾演练个数。

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDisasterRecoveryDrillsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDisasterRecoveryDrillsResponse struct{}"
	}

	return strings.Join([]string{"ListDisasterRecoveryDrillsResponse", string(data)}, " ")
}
