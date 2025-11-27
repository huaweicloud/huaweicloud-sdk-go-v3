package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWebTamperProtectionConfigResponse Response Object
type UpdateWebTamperProtectionConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateWebTamperProtectionConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWebTamperProtectionConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateWebTamperProtectionConfigResponse", string(data)}, " ")
}
