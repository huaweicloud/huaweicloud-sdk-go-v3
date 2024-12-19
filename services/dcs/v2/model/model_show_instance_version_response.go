package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceVersionResponse Response Object
type ShowInstanceVersionResponse struct {

	// 当前实例内核小版本信息。
	EngineMinorVersion *string `json:"engine_minor_version,omitempty"`

	// DCS实例最新的内核小版本信息。
	LatestEngineMinorVersion *string `json:"latest_engine_minor_version,omitempty"`

	// 当前实例proxy代理节点版本信息。
	ProxyMinorVersion *string `json:"proxy_minor_version,omitempty"`

	// DCS实例最新的proxy代理节点版本信息。
	LatestProxyMinorVersion *string `json:"latest_proxy_minor_version,omitempty"`

	// 当前实例内核是否可以升级。
	EngineMinorVersionUpgradable *bool `json:"engine_minor_version_upgradable,omitempty"`

	// 当前实例proxy代理节点是否可以升级。
	ProxyMinorVersionUpgradable *bool `json:"proxy_minor_version_upgradable,omitempty"`
	HttpStatusCode              int   `json:"-"`
}

func (o ShowInstanceVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceVersionResponse", string(data)}, " ")
}
