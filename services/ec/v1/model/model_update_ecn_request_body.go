package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateEcnRequestBody struct {

	// 企业连接网络名字
	Name *string `json:"name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 企业连接网络AS号
	EcnAsn *int64 `json:"ecn_asn,omitempty"`

	// 智能企业网关AS号
	IegAsn *int64 `json:"ieg_asn,omitempty"`

	// 分支互联开关
	HubEnable *bool `json:"hub_enable,omitempty"`
}

func (o UpdateEcnRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEcnRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEcnRequestBody", string(data)}, " ")
}
