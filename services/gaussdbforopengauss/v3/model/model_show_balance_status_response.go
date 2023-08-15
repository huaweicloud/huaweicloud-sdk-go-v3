package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBalanceStatusResponse Response Object
type ShowBalanceStatusResponse struct {

	// 平衡状态。
	Balanced       *bool `json:"balanced,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowBalanceStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBalanceStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowBalanceStatusResponse", string(data)}, " ")
}
