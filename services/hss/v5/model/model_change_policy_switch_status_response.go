package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangePolicySwitchStatusResponse Response Object
type ChangePolicySwitchStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangePolicySwitchStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePolicySwitchStatusResponse struct{}"
	}

	return strings.Join([]string{"ChangePolicySwitchStatusResponse", string(data)}, " ")
}
