package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// CreateGlobalDcGatewayEntry global-dc-gateway详情。
type CreateGlobalDcGatewayEntry struct {

	// 专线全域接入网关（global-dc-gateway）ID
	Id *string `json:"id,omitempty"`

	// 项目ID。
	TenantId *string `json:"tenant_id,omitempty"`

	// global-dc-gateway名字。
	Name *string `json:"name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// global-dc-gateway所属的企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// DGW加载的全球中心网络实例的ID
	GlobalCenterNetworkId *string `json:"global_center_network_id,omitempty"`

	// DGW对应BGP的ASN编号
	BgpAsn *int64 `json:"bgp_asn,omitempty"`

	// DGW所属Region
	RegionId *string `json:"region_id,omitempty"`

	// DGW创建网关设备归属的位置
	LocationName *string `json:"location_name,omitempty"`

	// 全域接入网关(GDGW)上关联连接的数量，表示DGW挂载ER的数量
	CurrentPeerLinkCount *int32 `json:"current_peer_link_count,omitempty"`

	// 该全域接入网关上GDGW允许创建关联连接（PeerLink）的数量
	AvailablePeerLinkCount *int32 `json:"available_peer_link_count,omitempty"`

	// global-dc-gateway关联TAG。
	Tags *[]Tag `json:"tags,omitempty"`

	// 该GDGW的管理状态，true为激活状态、false为冻结状态
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Status *GlobalDcGatewayStatus `json:"status,omitempty"`

	// 创建时间。
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// 网关的地址簇，IPv4或者ipv6和IPv4双栈 - ipv4 - dual
	AddressFamily *CreateGlobalDcGatewayEntryAddressFamily `json:"address_family,omitempty"`
}

func (o CreateGlobalDcGatewayEntry) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalDcGatewayEntry struct{}"
	}

	return strings.Join([]string{"CreateGlobalDcGatewayEntry", string(data)}, " ")
}

type CreateGlobalDcGatewayEntryAddressFamily struct {
	value string
}

type CreateGlobalDcGatewayEntryAddressFamilyEnum struct {
	IPV4 CreateGlobalDcGatewayEntryAddressFamily
	DUAL CreateGlobalDcGatewayEntryAddressFamily
}

func GetCreateGlobalDcGatewayEntryAddressFamilyEnum() CreateGlobalDcGatewayEntryAddressFamilyEnum {
	return CreateGlobalDcGatewayEntryAddressFamilyEnum{
		IPV4: CreateGlobalDcGatewayEntryAddressFamily{
			value: "ipv4",
		},
		DUAL: CreateGlobalDcGatewayEntryAddressFamily{
			value: "dual",
		},
	}
}

func (c CreateGlobalDcGatewayEntryAddressFamily) Value() string {
	return c.value
}

func (c CreateGlobalDcGatewayEntryAddressFamily) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateGlobalDcGatewayEntryAddressFamily) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
