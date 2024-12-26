package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoProtectStatusResponse Response Object
type ShowAutoProtectStatusResponse struct {
	Data           *FirewallStatusVo `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowAutoProtectStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoProtectStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoProtectStatusResponse", string(data)}, " ")
}
