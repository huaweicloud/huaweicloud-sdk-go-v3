package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAdminHealthCheckResponse Response Object
type ShowAdminHealthCheckResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ShowAdminHealthCheckResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAdminHealthCheckResponse struct{}"
	}

	return strings.Join([]string{"ShowAdminHealthCheckResponse", string(data)}, " ")
}
