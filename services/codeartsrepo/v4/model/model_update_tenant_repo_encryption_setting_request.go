package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTenantRepoEncryptionSettingRequest Request Object
type UpdateTenantRepoEncryptionSettingRequest struct {

	// **参数解释：** 租户id
	TenantId string `json:"tenant_id"`

	Body *TenantRepoEncryptionSettingRequestDto `json:"body,omitempty"`
}

func (o UpdateTenantRepoEncryptionSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTenantRepoEncryptionSettingRequest struct{}"
	}

	return strings.Join([]string{"UpdateTenantRepoEncryptionSettingRequest", string(data)}, " ")
}
