package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CadDomainSwitchRequest struct {

	// 域名id
	DomainId string `json:"domain_id"`

	// 是否开启WEB基础防护开关。0 - 开启， 1 - 关闭
	WafSwitch int32 `json:"waf_switch"`

	// 是否开启CC防护开关。0 - 开启， 1 - 关闭。开启CC开关必须同时开启WEB基础防护开关
	CcSwitch int32 `json:"cc_switch"`
}

func (o CadDomainSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CadDomainSwitchRequest struct{}"
	}

	return strings.Join([]string{"CadDomainSwitchRequest", string(data)}, " ")
}
