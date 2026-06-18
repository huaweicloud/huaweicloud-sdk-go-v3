package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTenantRepoEncryptionSettingRequest Request Object
type ShowTenantRepoEncryptionSettingRequest struct {

	// **参数解释：** 租户id
	TenantId string `json:"tenant_id"`
}

func (o ShowTenantRepoEncryptionSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTenantRepoEncryptionSettingRequest struct{}"
	}

	return strings.Join([]string{"ShowTenantRepoEncryptionSettingRequest", string(data)}, " ")
}
