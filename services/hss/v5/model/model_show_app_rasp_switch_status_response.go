package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppRaspSwitchStatusResponse Response Object
type ShowAppRaspSwitchStatusResponse struct {

	// 开启状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAppRaspSwitchStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppRaspSwitchStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowAppRaspSwitchStatusResponse", string(data)}, " ")
}
