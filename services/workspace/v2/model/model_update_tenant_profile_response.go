package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTenantProfileResponse Response Object
type UpdateTenantProfileResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTenantProfileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTenantProfileResponse struct{}"
	}

	return strings.Join([]string{"UpdateTenantProfileResponse", string(data)}, " ")
}
