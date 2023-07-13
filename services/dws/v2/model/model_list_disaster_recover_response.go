package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDisasterRecoverResponse Response Object
type ListDisasterRecoverResponse struct {

	// 容灾对象
	DisasterRecovery *[]DisasterRecovery `json:"disaster_recovery,omitempty"`
	HttpStatusCode   int                 `json:"-"`
}

func (o ListDisasterRecoverResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDisasterRecoverResponse struct{}"
	}

	return strings.Join([]string{"ListDisasterRecoverResponse", string(data)}, " ")
}
