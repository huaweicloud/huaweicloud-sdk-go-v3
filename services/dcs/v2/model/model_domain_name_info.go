package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DomainNameInfo struct {

	// 是否开启公网域名解析。 - true：开启 - false：未开启
	SupportPublicResolve *bool `json:"support_public_resolve,omitempty"`

	// 当前域名是否已为最新。 - true：是 - false：否
	IsLatestRules *bool `json:"is_latest_rules,omitempty"`

	// 域名的区域后缀。
	ZoneName *string `json:"zone_name,omitempty"`

	// 历史域名信息。
	HistoryDomainNames *[]DomainNameEntity `json:"history_domain_names,omitempty"`
}

func (o DomainNameInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainNameInfo struct{}"
	}

	return strings.Join([]string{"DomainNameInfo", string(data)}, " ")
}
