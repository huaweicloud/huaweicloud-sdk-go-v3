package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTenantDurationCfgRequest Request Object
type ShowTenantDurationCfgRequest struct {
}

func (o ShowTenantDurationCfgRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTenantDurationCfgRequest struct{}"
	}

	return strings.Join([]string{"ShowTenantDurationCfgRequest", string(data)}, " ")
}
