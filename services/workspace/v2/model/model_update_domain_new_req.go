package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainNewReq 更新域控的配置信息请求。
type UpdateDomainNewReq struct {
	UosDomainInfo *UpdateUosDomainInfo `json:"uos_domain_info,omitempty"`

	AdDomainInfo *AdDomain `json:"ad_domain_info,omitempty"`

	AuthType *DomainType `json:"auth_type"`

	// 认证配置id。
	AuthConfigId *string `json:"auth_config_id,omitempty"`
}

func (o UpdateDomainNewReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainNewReq struct{}"
	}

	return strings.Join([]string{"UpdateDomainNewReq", string(data)}, " ")
}
