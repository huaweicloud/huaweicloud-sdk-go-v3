package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTenantDevelopModeRequest Request Object
type ShowTenantDevelopModeRequest struct {
}

func (o ShowTenantDevelopModeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTenantDevelopModeRequest struct{}"
	}

	return strings.Join([]string{"ShowTenantDevelopModeRequest", string(data)}, " ")
}
