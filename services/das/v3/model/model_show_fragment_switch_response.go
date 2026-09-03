package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFragmentSwitchResponse Response Object
type ShowFragmentSwitchResponse struct {

	// 开关名称
	SwitchName *string `json:"switch_name,omitempty"`

	// 是否开启
	Open           *bool `json:"open,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowFragmentSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFragmentSwitchResponse struct{}"
	}

	return strings.Join([]string{"ShowFragmentSwitchResponse", string(data)}, " ")
}
