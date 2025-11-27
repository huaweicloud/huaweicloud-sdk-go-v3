package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProviderAgencyUrnPrimitiveTypeHolder struct {

	// 自定义provider所绑定的IAM委托URN，provider_agency_name和provider_agency_urn最多只能提供一个。
	ProviderAgencyUrn *string `json:"provider_agency_urn,omitempty"`
}

func (o ProviderAgencyUrnPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProviderAgencyUrnPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"ProviderAgencyUrnPrimitiveTypeHolder", string(data)}, " ")
}
