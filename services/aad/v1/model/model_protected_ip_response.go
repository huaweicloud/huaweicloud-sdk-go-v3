package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ProtectedIpResponse 防护ip列表
type ProtectedIpResponse struct {

	// 防护IP的Id
	Id string `json:"id"`

	// 防护IP
	Ip string `json:"ip"`

	// 类型。VPN：VPN；NAT：NAT；VIP：VIP；CCI：CCI；EIP：弹性公网IP；ELB：弹性负载均衡；REROUTING_IP：REROUTING_IP；SubEni：SubEni；NetInterFace：NetInterFace；
	Type ProtectedIpResponseType `json:"type"`

	// 名字
	Name *string `json:"name,omitempty"`

	// 状态：0 - 正常， 1 - 清洗中， 2 - 黑洞中
	Status int32 `json:"status"`

	StatusDetail *IpStatusDetail `json:"status_detail,omitempty"`

	// 策略名
	PolicyName string `json:"policy_name"`

	// 所属region
	Region string `json:"region"`

	// 防护包id
	PackageId string `json:"package_id"`

	// 防护包名
	PackageName string `json:"package_name"`

	// TMS标签
	Tags *string `json:"tags,omitempty"`

	// 本地标签
	Tag *string `json:"tag,omitempty"`

	// 默认false，表示是否转售版的IP，不需要展示策略和报表
	IsResale bool `json:"is_resale"`

	// package_version。cnad_pro：专业版；cnad_ip：标准版；cnad_ep：铂金版；cnad_full_high：全力防高级版；cnad_vic：按需版；cnad_intl_ep：国际站铂金版
	PackageVersion ProtectedIpResponsePackageVersion `json:"package_version"`
}

func (o ProtectedIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectedIpResponse struct{}"
	}

	return strings.Join([]string{"ProtectedIpResponse", string(data)}, " ")
}

type ProtectedIpResponseType struct {
	value string
}

type ProtectedIpResponseTypeEnum struct {
	VPN            ProtectedIpResponseType
	NAT            ProtectedIpResponseType
	VIP            ProtectedIpResponseType
	CCI            ProtectedIpResponseType
	EIP            ProtectedIpResponseType
	ELB            ProtectedIpResponseType
	REROUTING_IP   ProtectedIpResponseType
	SUB_ENI        ProtectedIpResponseType
	NET_INTER_FACE ProtectedIpResponseType
}

func GetProtectedIpResponseTypeEnum() ProtectedIpResponseTypeEnum {
	return ProtectedIpResponseTypeEnum{
		VPN: ProtectedIpResponseType{
			value: "VPN",
		},
		NAT: ProtectedIpResponseType{
			value: "NAT",
		},
		VIP: ProtectedIpResponseType{
			value: "VIP",
		},
		CCI: ProtectedIpResponseType{
			value: "CCI",
		},
		EIP: ProtectedIpResponseType{
			value: "EIP",
		},
		ELB: ProtectedIpResponseType{
			value: "ELB",
		},
		REROUTING_IP: ProtectedIpResponseType{
			value: "REROUTING_IP",
		},
		SUB_ENI: ProtectedIpResponseType{
			value: "SubEni",
		},
		NET_INTER_FACE: ProtectedIpResponseType{
			value: "NetInterFace",
		},
	}
}

func (c ProtectedIpResponseType) Value() string {
	return c.value
}

func (c ProtectedIpResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProtectedIpResponseType) UnmarshalJSON(b []byte) error {
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

type ProtectedIpResponsePackageVersion struct {
	value string
}

type ProtectedIpResponsePackageVersionEnum struct {
	CNAD_PRO       ProtectedIpResponsePackageVersion
	CNAD_IP        ProtectedIpResponsePackageVersion
	CNAD_EP        ProtectedIpResponsePackageVersion
	CNAD_FULL_HIGH ProtectedIpResponsePackageVersion
	CNAD_VIC       ProtectedIpResponsePackageVersion
	CNAD_INTL_EP   ProtectedIpResponsePackageVersion
}

func GetProtectedIpResponsePackageVersionEnum() ProtectedIpResponsePackageVersionEnum {
	return ProtectedIpResponsePackageVersionEnum{
		CNAD_PRO: ProtectedIpResponsePackageVersion{
			value: "cnad_pro",
		},
		CNAD_IP: ProtectedIpResponsePackageVersion{
			value: "cnad_ip",
		},
		CNAD_EP: ProtectedIpResponsePackageVersion{
			value: "cnad_ep",
		},
		CNAD_FULL_HIGH: ProtectedIpResponsePackageVersion{
			value: "cnad_full_high",
		},
		CNAD_VIC: ProtectedIpResponsePackageVersion{
			value: "cnad_vic",
		},
		CNAD_INTL_EP: ProtectedIpResponsePackageVersion{
			value: "cnad_intl_ep",
		},
	}
}

func (c ProtectedIpResponsePackageVersion) Value() string {
	return c.value
}

func (c ProtectedIpResponsePackageVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProtectedIpResponsePackageVersion) UnmarshalJSON(b []byte) error {
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
