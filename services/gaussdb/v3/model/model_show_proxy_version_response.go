package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProxyVersionResponse Response Object
type ShowProxyVersionResponse struct {

	// 当前代理版本
	CurrentVersion *string `json:"current_version,omitempty"`

	// 最新代理版本
	LatestVersion *string `json:"latest_version,omitempty"`

	// 是否能升级
	CanUpgrade     *bool `json:"can_upgrade,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowProxyVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProxyVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowProxyVersionResponse", string(data)}, " ")
}
