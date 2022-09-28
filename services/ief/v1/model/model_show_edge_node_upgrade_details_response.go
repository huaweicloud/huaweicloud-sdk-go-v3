package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowEdgeNodeUpgradeDetailsResponse struct {

	// 是否升级成功
	UpgradeEnable *bool `json:"upgrade_enable,omitempty"`

	// 未升级成功的原因，若升级成功，返回值为空字符串
	Reason *string `json:"reason,omitempty"`

	// 升级成功后，当前边缘节点的版本
	Version        *string `json:"version,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowEdgeNodeUpgradeDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEdgeNodeUpgradeDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowEdgeNodeUpgradeDetailsResponse", string(data)}, " ")
}
