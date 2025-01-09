package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateTenantConfigReq struct {
	FunctionConfig *FunctionConfig `json:"function_config,omitempty"`
}

func (o UpdateTenantConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTenantConfigReq struct{}"
	}

	return strings.Join([]string{"UpdateTenantConfigReq", string(data)}, " ")
}
