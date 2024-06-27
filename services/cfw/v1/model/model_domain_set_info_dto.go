package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DomainSetInfoDto struct {

	// 域名
	DomainName string `json:"domain_name"`

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
