package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProviderAgencyNamePrimitiveTypeHolder struct {

	// 自定义provider所绑定的IAM委托名称，provider_agency_name和provider_agency_urn最多只能提供一个。
	ProviderAgencyName *string `json:"provider_agency_name,omitempty"`
}

func (o ProviderAgencyNamePrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProviderAgencyNamePrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"ProviderAgencyNamePrimitiveTypeHolder", string(data)}, " ")
}
