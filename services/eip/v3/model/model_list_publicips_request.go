package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ListPublicipsRequest struct {
	// 分页查询起始的资源ID，为空时为查询第一页

	Marker *string `json:"marker,omitempty"`
	// 分页查询起始的资源序号

	Offset *int32 `json:"offset,omitempty"`
	// 每页返回的个数取值范围：0~[2000]，其中2000为局点差异项，具体取值由局点决定

	Limit *int32 `json:"limit,omitempty"`
	// 显示，形式为\"fields=id&fields=owner&...\"  支持字段：id/project_id/ip_version/type/public_ip_address/public_ipv6_address/network_type/status/description/created_at/updated_at/vnic/bandwidth/associate_instance_type/associate_instance_id/lock_status/billing_info/tags/enterprise_project_id/allow_share_bandwidth_types/public_border_group/alias/publicip_pool_name/publicip_pool_id

	Fields *[]string `json:"fields,omitempty"`
	// 排序，形式为\"sort_key=id\"  支持字段：id/public_ip_address/public_ipv6_address/ip_version/created_at/updated_at/public_border_group

	SortKey *ListPublicipsRequestSortKey `json:"sort_key,omitempty"`
	// 排序方向  取值范围：asc、desc

	SortDir *ListPublicipsRequestSortDir `json:"sort_dir,omitempty"`
	// 根据id过滤

	Id *[]string `json:"id,omitempty"`
	// 根据ip_version过滤  取值范围：4、6

	IpVersion *[]ListPublicipsRequestIpVersion `json:"ip_version,omitempty"`
	// 根据public_ip_address过滤

	PublicIpAddress *[]string `json:"public_ip_address,omitempty"`
	// 根据public_ip_address过滤，模糊搜索

	PublicIpAddressLike *string `json:"public_ip_address_like,omitempty"`
	// 根据public_ipv6_address过滤

	PublicIpv6Address *[]string `json:"public_ipv6_address,omitempty"`
	// 根据public_ipv6_address过滤，模糊搜索

	PublicIpv6AddressLike *string `json:"public_ipv6_address_like,omitempty"`
	// 根据type过滤  取值范围：EIP、DUALSTACK、DUALSTACK_SUBNET  EIP: 弹性公网IP   DUALSTACK: 双栈IPV6   DUALSTACK_SUBNET: 双栈子网

	Type *[]ListPublicipsRequestType `json:"type,omitempty"`
	// 根据network_type过滤  取值范围：5_telcom、5_union、5_bgp、5_sbgp、5_ipv6、5_graybgp

	NetworkType *[]ListPublicipsRequestNetworkType `json:"network_type,omitempty"`
	// 根据publicip_pool_name过滤  取值范围：5_telcom、5_union、5_bgp、5_sbgp、5_ipv6、5_graybgp、专属池名称等

	PublicipPoolName *[]string `json:"publicip_pool_name,omitempty"`
	// 根据status过滤  取值范围：FREEZED、DOWN、ACTIVE、ERROR

	Status *[]ListPublicipsRequestStatus `json:"status,omitempty"`
	// 根据alias模糊搜索

	AliasLike *string `json:"alias_like,omitempty"`
	// 根据alias过滤

	Alias *[]string `json:"alias,omitempty"`
	// 根据description过滤

	Description *[]string `json:"description,omitempty"`
	// 根据private_ip_address过滤

	VnicPrivateIpAddress *[]string `json:"vnic.private_ip_address,omitempty"`
	// 根据private_ip_address模糊搜索

	VnicPrivateIpAddressLike *string `json:"vnic.private_ip_address_like,omitempty"`
	// 根据device_id过滤

	VnicDeviceId *[]string `json:"vnic.device_id,omitempty"`
	// 根据device_owner过滤

	VnicDeviceOwner *[]string `json:"vnic.device_owner,omitempty"`
	// 根据vpc_id过滤

	VnicVpcId *[]string `json:"vnic.vpc_id,omitempty"`
	// 根据port_id过滤

	VnicPortId *[]string `json:"vnic.port_id,omitempty"`
	// 根据device_owner_prefixlike模糊搜索

	VnicDeviceOwnerPrefixlike *string `json:"vnic.device_owner_prefixlike,omitempty"`
	// 根据instance_type过滤

	VnicInstanceType *[]string `json:"vnic.instance_type,omitempty"`
	// 根据instance_id过滤

	VnicInstanceId *[]string `json:"vnic.instance_id,omitempty"`
	// 根据id过滤

	BandwidthId *[]string `json:"bandwidth.id,omitempty"`
	// 根据name过滤

	BandwidthName *[]string `json:"bandwidth.name,omitempty"`
	// 根据name模糊过滤

	BandwidthNameLike *[]string `json:"bandwidth.name_like,omitempty"`
	// 根据size过滤

	BandwidthSize *[]int32 `json:"bandwidth.size,omitempty"`
	// 根据share_type过滤

	BandwidthShareType *[]ListPublicipsRequestBandwidthShareType `json:"bandwidth.share_type,omitempty"`
	// 根据charge_mode过滤

	BandwidthChargeMode *[]ListPublicipsRequestBandwidthChargeMode `json:"bandwidth.charge_mode,omitempty"`
	// 根据billing_info过滤

	BillingInfo *[]string `json:"billing_info,omitempty"`
	// 根据订单模式过滤,   取值范围：YEARLY_MONTHLY、PAY_PER_USE

	BillingMode *ListPublicipsRequestBillingMode `json:"billing_mode,omitempty"`
	// 根据associate_instance_type过滤  取值范围：PORT、NATGW、ELB、VPN、ELBV1

	AssociateInstanceType *[]ListPublicipsRequestAssociateInstanceType `json:"associate_instance_type,omitempty"`
	// 根据associate_instance_id过滤

	AssociateInstanceId *[]string `json:"associate_instance_id,omitempty"`
	// 根据enterprise_project_id过滤

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
	// 根据public_border_group过滤

	PublicBorderGroup *[]string `json:"public_border_group,omitempty"`
	// 共享带宽类型，根据任一共享带宽类型过滤EIP列表。 可以指定多个带宽类型，不同的带宽类型间用逗号分隔。

	AllowShareBandwidthTypeAny *[]string `json:"allow_share_bandwidth_type_any,omitempty"`
}

func (o ListPublicipsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicipsRequest struct{}"
	}

	return strings.Join([]string{"ListPublicipsRequest", string(data)}, " ")
}

type ListPublicipsRequestSortKey struct {
	value string
}

type ListPublicipsRequestSortKeyEnum struct {
	ID                  ListPublicipsRequestSortKey
	PUBLIC_IP_ADDRESS   ListPublicipsRequestSortKey
	PUBLIC_IPV6_ADDRESS ListPublicipsRequestSortKey
	IP_VERSION          ListPublicipsRequestSortKey
	CREATED_AT          ListPublicipsRequestSortKey
	UPDATED_AT          ListPublicipsRequestSortKey
	PUBLIC_BORDER_GROUP ListPublicipsRequestSortKey
}

func GetListPublicipsRequestSortKeyEnum() ListPublicipsRequestSortKeyEnum {
	return ListPublicipsRequestSortKeyEnum{
		ID: ListPublicipsRequestSortKey{
			value: "id",
		},
		PUBLIC_IP_ADDRESS: ListPublicipsRequestSortKey{
			value: "public_ip_address",
		},
		PUBLIC_IPV6_ADDRESS: ListPublicipsRequestSortKey{
			value: "public_ipv6_address",
		},
		IP_VERSION: ListPublicipsRequestSortKey{
			value: "ip_version",
		},
		CREATED_AT: ListPublicipsRequestSortKey{
			value: "created_at",
		},
		UPDATED_AT: ListPublicipsRequestSortKey{
			value: "updated_at",
		},
		PUBLIC_BORDER_GROUP: ListPublicipsRequestSortKey{
			value: "public_border_group",
		},
	}
}

func (c ListPublicipsRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPublicipsRequestSortKey) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListPublicipsRequestSortDir struct {
	value string
}

type ListPublicipsRequestSortDirEnum struct {
	ASC  ListPublicipsRequestSortDir
	DESC ListPublicipsRequestSortDir
}

func GetListPublicipsRequestSortDirEnum() ListPublicipsRequestSortDirEnum {
	return ListPublicipsRequestSortDirEnum{
		ASC: ListPublicipsRequestSortDir{
			value: "asc",
		},
		DESC: ListPublicipsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListPublicipsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPublicipsRequestSortDir) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListPublicipsRequestIpVersion struct {
	value int32
}

type ListPublicipsRequestIpVersionEnum struct {
	E_4 ListPublicipsRequestIpVersion
	E_6 ListPublicipsRequestIpVersion
}

func GetListPublicipsRequestIpVersionEnum() ListPublicipsRequestIpVersionEnum {
	return ListPublicipsRequestIpVersionEnum{
		E_4: ListPublicipsRequestIpVersion{
			value: 4,
		}, E_6: ListPublicipsRequestIpVersion{
			value: 6,
		},
	}
}

func (c ListPublicipsRequestIpVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPublicipsRequestIpVersion) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type ListPublicipsRequestType struct {
	value string
}

type ListPublicipsRequestTypeEnum struct {
	EIP              ListPublicipsRequestType
	DUALSTACK        ListPublicipsRequestType
	DUALSTACK_SUBNET ListPublicipsRequestType
}

func GetListPublicipsRequestTypeEnum() ListPublicipsRequestTypeEnum {
	return ListPublicipsRequestTypeEnum{
		EIP: ListPublicipsRequestType{
			value: "EIP",
		},
		DUALSTACK: ListPublicipsRequestType{
			value: "DUALSTACK",
		},
		DUALSTACK_SUBNET: ListPublicipsRequestType{
			value: "DUALSTACK_SUBNET",
		},
	}
}

func (c ListPublicipsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPublicipsRequestType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListPublicipsRequestNetworkType struct {
	value string
}

type ListPublicipsRequestNetworkTypeEnum struct {
	E_5_TELCOM  ListPublicipsRequestNetworkType
	E_5_UNION   ListPublicipsRequestNetworkType
	E_5_BGP     ListPublicipsRequestNetworkType
	E_5_SBGP    ListPublicipsRequestNetworkType
	E_5_IPV6    ListPublicipsRequestNetworkType
	E_5_GRAYBGP ListPublicipsRequestNetworkType
}

func GetListPublicipsRequestNetworkTypeEnum() ListPublicipsRequestNetworkTypeEnum {
	return ListPublicipsRequestNetworkTypeEnum{
		E_5_TELCOM: ListPublicipsRequestNetworkType{
			value: "5_telcom",
		},
		E_5_UNION: ListPublicipsRequestNetworkType{
			value: "5_union",
		},
		E_5_BGP: ListPublicipsRequestNetworkType{
			value: "5_bgp",
		},
		E_5_SBGP: ListPublicipsRequestNetworkType{
			value: "5_sbgp",
		},
		E_5_IPV6: ListPublicipsRequestNetworkType{
			value: "5_ipv6",
		},
		E_5_GRAYBGP: ListPublicipsRequestNetworkType{
			value: "5_graybgp",
		},
	}
}

func (c ListPublicipsRequestNetworkType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPublicipsRequestNetworkType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListPublicipsRequestStatus struct {
	value string
}

type ListPublicipsRequestStatusEnum struct {
	FREEZED ListPublicipsRequestStatus
	DOWN    ListPublicipsRequestStatus
	ACTIVE  ListPublicipsRequestStatus
	ERROR   ListPublicipsRequestStatus
}

func GetListPublicipsRequestStatusEnum() ListPublicipsRequestStatusEnum {
	return ListPublicipsRequestStatusEnum{
		FREEZED: ListPublicipsRequestStatus{
			value: "FREEZED",
		},
		DOWN: ListPublicipsRequestStatus{
			value: "DOWN",
		},
		ACTIVE: ListPublicipsRequestStatus{
			value: "ACTIVE",
		},
		ERROR: ListPublicipsRequestStatus{
			value: "ERROR",
		},
	}
}

func (c ListPublicipsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPublicipsRequestStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListPublicipsRequestBandwidthShareType struct {
	value string
}

type ListPublicipsRequestBandwidthShareTypeEnum struct {
	PER   ListPublicipsRequestBandwidthShareType
	WHOLE ListPublicipsRequestBandwidthShareType
}

func GetListPublicipsRequestBandwidthShareTypeEnum() ListPublicipsRequestBandwidthShareTypeEnum {
	return ListPublicipsRequestBandwidthShareTypeEnum{
		PER: ListPublicipsRequestBandwidthShareType{
			value: "PER",
		},
		WHOLE: ListPublicipsRequestBandwidthShareType{
			value: "WHOLE",
		},
	}
}

func (c ListPublicipsRequestBandwidthShareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPublicipsRequestBandwidthShareType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListPublicipsRequestBandwidthChargeMode struct {
	value string
}

type ListPublicipsRequestBandwidthChargeModeEnum struct {
	BANDWIDTH     ListPublicipsRequestBandwidthChargeMode
	TRAFFIC       ListPublicipsRequestBandwidthChargeMode
	E_95PEAK_PLUS ListPublicipsRequestBandwidthChargeMode
}

func GetListPublicipsRequestBandwidthChargeModeEnum() ListPublicipsRequestBandwidthChargeModeEnum {
	return ListPublicipsRequestBandwidthChargeModeEnum{
		BANDWIDTH: ListPublicipsRequestBandwidthChargeMode{
			value: "bandwidth",
		},
		TRAFFIC: ListPublicipsRequestBandwidthChargeMode{
			value: "traffic",
		},
		E_95PEAK_PLUS: ListPublicipsRequestBandwidthChargeMode{
			value: "95peak_plus",
		},
	}
}

func (c ListPublicipsRequestBandwidthChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPublicipsRequestBandwidthChargeMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListPublicipsRequestBillingMode struct {
	value string
}

type ListPublicipsRequestBillingModeEnum struct {
	YEARLY_MONTHLY ListPublicipsRequestBillingMode
	PAY_PER_USE    ListPublicipsRequestBillingMode
}

func GetListPublicipsRequestBillingModeEnum() ListPublicipsRequestBillingModeEnum {
	return ListPublicipsRequestBillingModeEnum{
		YEARLY_MONTHLY: ListPublicipsRequestBillingMode{
			value: "YEARLY_MONTHLY",
		},
		PAY_PER_USE: ListPublicipsRequestBillingMode{
			value: "PAY_PER_USE",
		},
	}
}

func (c ListPublicipsRequestBillingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPublicipsRequestBillingMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListPublicipsRequestAssociateInstanceType struct {
	value string
}

type ListPublicipsRequestAssociateInstanceTypeEnum struct {
	PORT  ListPublicipsRequestAssociateInstanceType
	NATGW ListPublicipsRequestAssociateInstanceType
	ELB   ListPublicipsRequestAssociateInstanceType
	VPN   ListPublicipsRequestAssociateInstanceType
	ELBV1 ListPublicipsRequestAssociateInstanceType
}

func GetListPublicipsRequestAssociateInstanceTypeEnum() ListPublicipsRequestAssociateInstanceTypeEnum {
	return ListPublicipsRequestAssociateInstanceTypeEnum{
		PORT: ListPublicipsRequestAssociateInstanceType{
			value: "PORT",
		},
		NATGW: ListPublicipsRequestAssociateInstanceType{
			value: "NATGW",
		},
		ELB: ListPublicipsRequestAssociateInstanceType{
			value: "ELB",
		},
		VPN: ListPublicipsRequestAssociateInstanceType{
			value: "VPN",
		},
		ELBV1: ListPublicipsRequestAssociateInstanceType{
			value: "ELBV1",
		},
	}
}

func (c ListPublicipsRequestAssociateInstanceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPublicipsRequestAssociateInstanceType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
