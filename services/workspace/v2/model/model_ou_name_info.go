package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OuNameInfo struct {

	// ouid。
	Id *string `json:"id,omitempty"`

	// 域id
	DomainId *string `json:"domain_id,omitempty"`

	// 域名称
	Domain *string `json:"domain,omitempty"`

	// OU名称
	OuName *string `json:"ou_name,omitempty"`

	// ouDn
	OuDn *string `json:"ou_dn,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o OuNameInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OuNameInfo struct{}"
	}

	return strings.Join([]string{"OuNameInfo", string(data)}, " ")
}
