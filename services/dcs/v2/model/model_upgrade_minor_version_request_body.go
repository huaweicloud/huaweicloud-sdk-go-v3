package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeMinorVersionRequestBody 升级实例小版本请求体，proxy_minor_version和engine_minor_version不能同时为空
type UpgradeMinorVersionRequestBody struct {

	// Proxy代理节点目标版本号，设置为latest时，即升级到最新版本。
	ProxyMinorVersion *string `json:"proxy_minor_version,omitempty"`

	// 引擎目标小版本号，设置为latest时，即升级到最新版本。
	EngineMinorVersion *string `json:"engine_minor_version,omitempty"`
}

func (o UpgradeMinorVersionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeMinorVersionRequestBody struct{}"
	}

	return strings.Join([]string{"UpgradeMinorVersionRequestBody", string(data)}, " ")
}
