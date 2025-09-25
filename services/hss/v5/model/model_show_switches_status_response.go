package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSwitchesStatusResponse Response Object
type ShowSwitchesStatusResponse struct {

	// 是否开启
	Enabled        *bool `json:"enabled,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowSwitchesStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSwitchesStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowSwitchesStatusResponse", string(data)}, " ")
}
