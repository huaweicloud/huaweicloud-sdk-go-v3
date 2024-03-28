package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SlDomainAccessSetting struct {

	// 设置调试域名是否可以访问，true为可以访问，false为禁止访问
	SlDomainAccessEnabled bool `json:"sl_domain_access_enabled"`
}

func (o SlDomainAccessSetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlDomainAccessSetting struct{}"
	}

	return strings.Join([]string{"SlDomainAccessSetting", string(data)}, " ")
}
