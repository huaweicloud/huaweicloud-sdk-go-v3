package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 托管专线对象信息
type HostedDirectConnect struct {

	// 托管专线ID
	Id *string `json:"id,omitempty"`

	// 实例所属项目ID。
	TenantId *string `json:"tenant_id,omitempty"`

	// 物理专线名字
	Name *string `json:"name,omitempty"`

	// 物理专线的描述信息
	Description *string `json:"description,omitempty"`

	// 物理专线接入带宽，单位Mbps。
	Bandwidth *int32 `json:"bandwidth,omitempty"`

	// 专线的接入位置信息
	Location *string `json:"location,omitempty"`

	// 物理专线对端所在的物理位置，省/市/街道或IDC名字。
	PeerLocation *string `json:"peer_location,omitempty"`

	// hosted物理专线对应的hosting物理专线的ID
	HostingId *string `json:"hosting_id,omitempty"`

	// 专线线路的提供商
	Provider *string `json:"provider,omitempty"`

	// 管理状态：true或false
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// hosted物理专线预分配的vlan。
	Vlan *int32 `json:"vlan,omitempty"`

	// 操作状态，合法值是：ACTIVE， ERROR，PENDING_CREATE，PENDING_UPDATE。ACTIVE：虚拟网关正常ERROR： 虚拟网关异常PENDING_CREATE：创建中PENDING_UPDATE：更新中
	Status *HostedDirectConnectStatus `json:"status,omitempty"`

	// 物理专线申请时间
	ApplyTime *string `json:"apply_time,omitempty"`

	// 物理专线创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 物理专线的运营商操作状态，合法值是：ACTIVE， DOWN
	ProviderStatus *HostedDirectConnectProviderStatus `json:"provider_status,omitempty"`
}

func (o HostedDirectConnect) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostedDirectConnect struct{}"
	}

	return strings.Join([]string{"HostedDirectConnect", string(data)}, " ")
}

type HostedDirectConnectStatus struct {
	value string
}

type HostedDirectConnectStatusEnum struct {
	BUILD          HostedDirectConnectStatus
	PAID           HostedDirectConnectStatus
	APPLY          HostedDirectConnectStatus
	PENDING_SURVEY HostedDirectConnectStatus
	ACTIVE         HostedDirectConnectStatus
	DOWN           HostedDirectConnectStatus
	ERROR          HostedDirectConnectStatus
	PENDING_DELETE HostedDirectConnectStatus
	DELETED        HostedDirectConnectStatus
	DENY           HostedDirectConnectStatus
	PENDING_PAY    HostedDirectConnectStatus
	ORDERING       HostedDirectConnectStatus
	ACCEPT         HostedDirectConnectStatus
	REJECTED       HostedDirectConnectStatus
}

func GetHostedDirectConnectStatusEnum() HostedDirectConnectStatusEnum {
	return HostedDirectConnectStatusEnum{
		BUILD: HostedDirectConnectStatus{
			value: "BUILD",
		},
		PAID: HostedDirectConnectStatus{
			value: "PAID",
		},
		APPLY: HostedDirectConnectStatus{
			value: "APPLY",
		},
		PENDING_SURVEY: HostedDirectConnectStatus{
			value: "PENDING_SURVEY",
		},
		ACTIVE: HostedDirectConnectStatus{
			value: "ACTIVE",
		},
		DOWN: HostedDirectConnectStatus{
			value: "DOWN",
		},
		ERROR: HostedDirectConnectStatus{
			value: "ERROR",
		},
		PENDING_DELETE: HostedDirectConnectStatus{
			value: "PENDING_DELETE",
		},
		DELETED: HostedDirectConnectStatus{
			value: "DELETED",
		},
		DENY: HostedDirectConnectStatus{
			value: "DENY",
		},
		PENDING_PAY: HostedDirectConnectStatus{
			value: "PENDING_PAY",
		},
		ORDERING: HostedDirectConnectStatus{
			value: "ORDERING",
		},
		ACCEPT: HostedDirectConnectStatus{
			value: "ACCEPT",
		},
		REJECTED: HostedDirectConnectStatus{
			value: "REJECTED",
		},
	}
}

func (c HostedDirectConnectStatus) Value() string {
	return c.value
}

func (c HostedDirectConnectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HostedDirectConnectStatus) UnmarshalJSON(b []byte) error {
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

type HostedDirectConnectProviderStatus struct {
	value string
}

type HostedDirectConnectProviderStatusEnum struct {
	ACTIVE HostedDirectConnectProviderStatus
	DOWN   HostedDirectConnectProviderStatus
}

func GetHostedDirectConnectProviderStatusEnum() HostedDirectConnectProviderStatusEnum {
	return HostedDirectConnectProviderStatusEnum{
		ACTIVE: HostedDirectConnectProviderStatus{
			value: "ACTIVE",
		},
		DOWN: HostedDirectConnectProviderStatus{
			value: "DOWN",
		},
	}
}

func (c HostedDirectConnectProviderStatus) Value() string {
	return c.value
}

func (c HostedDirectConnectProviderStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HostedDirectConnectProviderStatus) UnmarshalJSON(b []byte) error {
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
