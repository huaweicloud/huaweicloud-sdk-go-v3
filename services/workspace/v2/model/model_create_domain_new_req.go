package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDomainNewReq 配置域控的配置信息请求。
type CreateDomainNewReq struct {
	UosDomainInfo *CreateUosDomainInfo `json:"uos_domain_info,omitempty"`

	AdDomainInfo *AdDomain `json:"ad_domain_info,omitempty"`

	AuthType *DomainType `json:"auth_type,omitempty"`
}

func (o CreateDomainNewReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainNewReq struct{}"
	}

	return strings.Join([]string{"CreateDomainNewReq", string(data)}, " ")
}
