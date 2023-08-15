package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowAlertRspUserInfo struct {

	// Id value
	UserId *string `json:"user_id,omitempty"`

	// The name, display only
	UserName *string `json:"user_name,omitempty"`
}

func (o ShowAlertRspUserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlertRspUserInfo struct{}"
	}

	return strings.Join([]string{"ShowAlertRspUserInfo", string(data)}, " ")
}
