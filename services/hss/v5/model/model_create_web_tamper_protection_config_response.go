package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWebTamperProtectionConfigResponse Response Object
type CreateWebTamperProtectionConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateWebTamperProtectionConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWebTamperProtectionConfigResponse struct{}"
	}

	return strings.Join([]string{"CreateWebTamperProtectionConfigResponse", string(data)}, " ")
}
