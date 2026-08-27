package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTenantVersionConfigRequest Request Object
type ShowTenantVersionConfigRequest struct {

	// 版本配置ID
	VersionConfigId string `json:"version_config_id"`
}

func (o ShowTenantVersionConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTenantVersionConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowTenantVersionConfigRequest", string(data)}, " ")
}
