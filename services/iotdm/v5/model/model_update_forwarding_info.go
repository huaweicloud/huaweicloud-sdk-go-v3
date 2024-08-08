package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateForwardingInfo 企业版实例的SNAT配置，配置开启后，企业版实例可以在公共网络中进行外部通信。 约束：只有企业版实例支持修改流转配置。
type UpdateForwardingInfo struct {

	// **参数说明**：是否启用SNAT配置。企业版实例开启SNAT配置后，可以在公共网络中进行外部通信。 约束：只有企业版实例支持配置SNAT配置，SNAT配置开启后将不支持关闭 **取值范围**： - true: 启用SNAT配置
	EnableSnat bool `json:"enable_snat"`
}

func (o UpdateForwardingInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateForwardingInfo struct{}"
	}

	return strings.Join([]string{"UpdateForwardingInfo", string(data)}, " ")
}
