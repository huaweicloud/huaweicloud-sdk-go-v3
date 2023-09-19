package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAdminConfigRequest Request Object
type ShowAdminConfigRequest struct {
}

func (o ShowAdminConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAdminConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowAdminConfigRequest", string(data)}, " ")
}
