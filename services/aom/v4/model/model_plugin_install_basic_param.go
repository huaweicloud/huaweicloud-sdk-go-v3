package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PluginInstallBasicParam struct {

	// 指定安装的ICAgent版本。
	InstallVersion *string `json:"install_version,omitempty"`

	// IAM账号AK，选填。.
	DomainAk *string `json:"domain_ak,omitempty"`

	// IAM账号SK，选填。
	DomainSk *string `json:"domain_sk,omitempty"`
}

func (o PluginInstallBasicParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginInstallBasicParam struct{}"
	}

	return strings.Join([]string{"PluginInstallBasicParam", string(data)}, " ")
}
