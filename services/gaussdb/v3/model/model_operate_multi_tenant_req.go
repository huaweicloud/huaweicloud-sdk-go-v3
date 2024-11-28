package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OperateMultiTenantReq 开启/关闭多租特性请求体
type OperateMultiTenantReq struct {

	// 实例多租特性开关。 - true:开启 - false:关闭。
	MultiTenantSwitch bool `json:"multi_tenant_switch"`
}

func (o OperateMultiTenantReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateMultiTenantReq struct{}"
	}

	return strings.Join([]string{"OperateMultiTenantReq", string(data)}, " ")
}
