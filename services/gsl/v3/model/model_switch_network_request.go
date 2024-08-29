package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchNetworkRequest Request Object
type SwitchNetworkRequest struct {

	// sim卡id
	SimCardId int64 `json:"sim_card_id"`

	Body *NetworkSwitchReq `json:"body,omitempty"`
}

func (o SwitchNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchNetworkRequest struct{}"
	}

	return strings.Join([]string{"SwitchNetworkRequest", string(data)}, " ")
}
