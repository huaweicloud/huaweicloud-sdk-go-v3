package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProxyUpgradeProxyVersionRequest proxy升级内核版本请求体对象。
type ProxyUpgradeProxyVersionRequest struct {

	// 升级前源内核版本号
	SourceVersion string `json:"source_version"`

	// 目标升级内核版本号
	TargetVersion string `json:"target_version"`
}

func (o ProxyUpgradeProxyVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProxyUpgradeProxyVersionRequest struct{}"
	}

	return strings.Join([]string{"ProxyUpgradeProxyVersionRequest", string(data)}, " ")
}
