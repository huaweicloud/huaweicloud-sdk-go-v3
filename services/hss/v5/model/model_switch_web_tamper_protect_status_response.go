package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchWebTamperProtectStatusResponse Response Object
type SwitchWebTamperProtectStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchWebTamperProtectStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchWebTamperProtectStatusResponse struct{}"
	}

	return strings.Join([]string{"SwitchWebTamperProtectStatusResponse", string(data)}, " ")
}
