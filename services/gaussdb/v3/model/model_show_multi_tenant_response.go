package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMultiTenantResponse Response Object
type ShowMultiTenantResponse struct {

	// 实例多租特性开关。 - true:开启 - false:关闭。
	MultiTenantSwitch *bool `json:"multi_tenant_switch,omitempty"`
	HttpStatusCode    int   `json:"-"`
}

func (o ShowMultiTenantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMultiTenantResponse struct{}"
	}

	return strings.Join([]string{"ShowMultiTenantResponse", string(data)}, " ")
}
