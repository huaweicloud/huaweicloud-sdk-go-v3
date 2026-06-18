package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTenantDevelopModeResponse Response Object
type ShowTenantDevelopModeResponse struct {

	// **参数解释：** 是否开启cr模式。
	CrEnable *bool `json:"cr_enable,omitempty"`

	// **参数解释：** 是否开启租户下加密设置。
	RepoEncryptionEnabled *bool `json:"repo_encryption_enabled,omitempty"`
	HttpStatusCode        int   `json:"-"`
}

func (o ShowTenantDevelopModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTenantDevelopModeResponse struct{}"
	}

	return strings.Join([]string{"ShowTenantDevelopModeResponse", string(data)}, " ")
}
