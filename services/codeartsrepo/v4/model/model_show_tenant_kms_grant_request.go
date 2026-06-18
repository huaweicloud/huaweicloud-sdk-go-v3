package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTenantKmsGrantRequest Request Object
type ShowTenantKmsGrantRequest struct {

	// **参数解释：** 租户id
	TenantId string `json:"tenant_id"`
}

func (o ShowTenantKmsGrantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTenantKmsGrantRequest struct{}"
	}

	return strings.Join([]string{"ShowTenantKmsGrantRequest", string(data)}, " ")
}
