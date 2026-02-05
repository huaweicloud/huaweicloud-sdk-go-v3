package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AddEipV2Response Response Object
type AddEipV2Response struct {

	// 弹性公网IP编号
	EipId *string `json:"eip_id,omitempty"`

	// 弹性公网IP
	EipAddress *string `json:"eip_address,omitempty"`

	// 弹性公网IP状态。 - FREEZED：冻结 - BIND_ERROR：绑定失败 - BINDING：绑定中 - PENDING_DELETE：释放中 - PENDING_CREATE：创建中 - NOTIFYING：创建中 - NOTIFY_DELETE：释放中 - PENDING_UPDATE：更新中 - DOWN：未绑定 - ACTIVE：绑定 - ELB：绑定ELB - VPN：绑定VPN - ERROR：失败
	EipStatus *AddEipV2ResponseEipStatus `json:"eip_status,omitempty"`

	// 弹性公网IP(IPV6)
	EipIpv6Address *string `json:"eip_ipv6_address,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddEipV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddEipV2Response struct{}"
	}

	return strings.Join([]string{"AddEipV2Response", string(data)}, " ")
}

type AddEipV2ResponseEipStatus struct {
	value string
}

type AddEipV2ResponseEipStatusEnum struct {
	FREEZED        AddEipV2ResponseEipStatus
	BIND_ERROR     AddEipV2ResponseEipStatus
	BINDING        AddEipV2ResponseEipStatus
	PENDING_DELETE AddEipV2ResponseEipStatus
	PENDING_CREATE AddEipV2ResponseEipStatus
	NOTIFYING      AddEipV2ResponseEipStatus
	NOTIFY_DELETE  AddEipV2ResponseEipStatus
	PENDING_UPDATE AddEipV2ResponseEipStatus
	DOWN           AddEipV2ResponseEipStatus
	ACTIVE         AddEipV2ResponseEipStatus
	ELB            AddEipV2ResponseEipStatus
	VPN            AddEipV2ResponseEipStatus
	ERROR          AddEipV2ResponseEipStatus
}

func GetAddEipV2ResponseEipStatusEnum() AddEipV2ResponseEipStatusEnum {
	return AddEipV2ResponseEipStatusEnum{
		FREEZED: AddEipV2ResponseEipStatus{
			value: "FREEZED",
		},
		BIND_ERROR: AddEipV2ResponseEipStatus{
			value: "BIND_ERROR",
		},
		BINDING: AddEipV2ResponseEipStatus{
			value: "BINDING",
		},
		PENDING_DELETE: AddEipV2ResponseEipStatus{
			value: "PENDING_DELETE",
		},
		PENDING_CREATE: AddEipV2ResponseEipStatus{
			value: "PENDING_CREATE",
		},
		NOTIFYING: AddEipV2ResponseEipStatus{
			value: "NOTIFYING",
		},
		NOTIFY_DELETE: AddEipV2ResponseEipStatus{
			value: "NOTIFY_DELETE",
		},
		PENDING_UPDATE: AddEipV2ResponseEipStatus{
			value: "PENDING_UPDATE",
		},
		DOWN: AddEipV2ResponseEipStatus{
			value: "DOWN",
		},
		ACTIVE: AddEipV2ResponseEipStatus{
			value: "ACTIVE",
		},
		ELB: AddEipV2ResponseEipStatus{
			value: "ELB",
		},
		VPN: AddEipV2ResponseEipStatus{
			value: "VPN",
		},
		ERROR: AddEipV2ResponseEipStatus{
			value: "ERROR",
		},
	}
}

func (c AddEipV2ResponseEipStatus) Value() string {
	return c.value
}

func (c AddEipV2ResponseEipStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddEipV2ResponseEipStatus) UnmarshalJSON(b []byte) error {
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
