package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWebTamperProtectionConfigResponse Response Object
type DeleteWebTamperProtectionConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteWebTamperProtectionConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWebTamperProtectionConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteWebTamperProtectionConfigResponse", string(data)}, " ")
}
