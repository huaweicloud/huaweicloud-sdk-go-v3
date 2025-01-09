package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTenantProfileRequest Request Object
type UpdateTenantProfileRequest struct {

	// 修改租户功能开关。
	Body map[string]bool `json:"body,omitempty"`
}

func (o UpdateTenantProfileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTenantProfileRequest struct{}"
	}

	return strings.Join([]string{"UpdateTenantProfileRequest", string(data)}, " ")
}
