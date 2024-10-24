package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAdminHealthCheckRequest Request Object
type ShowAdminHealthCheckRequest struct {
}

func (o ShowAdminHealthCheckRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAdminHealthCheckRequest struct{}"
	}

	return strings.Join([]string{"ShowAdminHealthCheckRequest", string(data)}, " ")
}
