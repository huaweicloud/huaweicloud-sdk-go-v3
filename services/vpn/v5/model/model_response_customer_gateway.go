package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResponseCustomerGateway struct {

	// 网关的ID
	Id *string `json:"id,omitempty"`

	// 网关名称
	Name *string `json:"name,omitempty"`

	// 网关的bgp asn号
	BgpAsn *int64 `json:"bgp_asn,omitempty"`

	// 对端网关标识类型
	IdType *string `json:"id_type,omitempty"`

	// 对端网关标识值
	IdValue *string `json:"id_value,omitempty"`

	CaCertificate *CaCertificate `json:"ca_certificate,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 标签
	Tags *[]VpnResourceTag `json:"tags,omitempty"`
}

func (o ResponseCustomerGateway) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResponseCustomerGateway struct{}"
	}

	return strings.Join([]string{"ResponseCustomerGateway", string(data)}, " ")
}
