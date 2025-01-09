package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InitializeTenantReq 租户信息配置。
type InitializeTenantReq struct {

	// 服务状态 * `active` - 激活 * `inactive` - 未激活
	ServiceStatus string `json:"service_status"`
}

func (o InitializeTenantReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InitializeTenantReq struct{}"
	}

	return strings.Join([]string{"InitializeTenantReq", string(data)}, " ")
}
