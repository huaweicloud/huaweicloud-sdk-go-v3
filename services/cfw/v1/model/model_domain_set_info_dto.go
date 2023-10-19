package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DomainSetInfoDto struct {

	// 域名组id
	DomainSetId *string `json:"domain_set_id,omitempty"`

	// 域名
	DomainName *string `json:"domain_name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o DomainSetInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainSetInfoDto struct{}"
	}

	return strings.Join([]string{"DomainSetInfoDto", string(data)}, " ")
}
