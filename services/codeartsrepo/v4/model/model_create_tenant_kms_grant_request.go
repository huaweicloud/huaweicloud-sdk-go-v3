package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTenantKmsGrantRequest Request Object
type CreateTenantKmsGrantRequest struct {

	// **参数解释：** 租户id
	TenantId string `json:"tenant_id"`
}

func (o CreateTenantKmsGrantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTenantKmsGrantRequest struct{}"
	}

	return strings.Join([]string{"CreateTenantKmsGrantRequest", string(data)}, " ")
}
