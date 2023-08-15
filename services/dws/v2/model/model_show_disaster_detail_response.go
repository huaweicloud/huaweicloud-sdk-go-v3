package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDisasterDetailResponse Response Object
type ShowDisasterDetailResponse struct {
	DisasterRecovery *DisasterRecoveryQueryResp `json:"disaster_recovery,omitempty"`
	HttpStatusCode   int                        `json:"-"`
}

func (o ShowDisasterDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDisasterDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowDisasterDetailResponse", string(data)}, " ")
}
