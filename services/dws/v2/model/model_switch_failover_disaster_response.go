package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchFailoverDisasterResponse Response Object
type SwitchFailoverDisasterResponse struct {
	DisasterRecovery *DisasterRecoveryId `json:"disaster_recovery,omitempty"`
	HttpStatusCode   int                 `json:"-"`
}

func (o SwitchFailoverDisasterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchFailoverDisasterResponse struct{}"
	}

	return strings.Join([]string{"SwitchFailoverDisasterResponse", string(data)}, " ")
}
