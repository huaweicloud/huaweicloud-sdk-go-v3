package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTenantConfigResponse Response Object
type UpdateTenantConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTenantConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTenantConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateTenantConfigResponse", string(data)}, " ")
}
