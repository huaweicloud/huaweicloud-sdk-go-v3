package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreatePublicipOption 创建弹性公网IP的IP对象
type CreatePublicipOption struct {

	// 功能说明：希望申请到的弹性公网IP的地址，不指定时由系统自动分配  约束：必须为IP地址格式，且必须在可用地址池范围内
	IpAddress *string `json:"ip_address,omitempty"`

	// 功能说明：弹性公网IP的类型  取值范围：5_telcom（电信），5_union（联通），5_bgp（全动态BGP），5_sbgp（静态BGP），5_ipv6  东北-大连：5_telcom、5_union  华南-广州：5_bgp、5_sbgp  华东-上海二：5_bgp、5_sbgp  华北-北京一：5_bgp、5_sbgp、5_ipv6  亚太-香港：5_bgp  亚太-曼谷：5_bgp  亚太-新加坡：5_bgp  非洲-约翰内斯堡：5_bgp  西南-贵阳一：5_bgp、5_sbgp  华北-北京四：5_bgp、5_sbgp  约束：必须是系统具体支持的类型。  publicip_id为IPv4端口，所以\"publicip_type\"字段未给定时，默认为5_bgp。
	Type string `json:"type"`

	// 功能说明：弹性IP弹性公网IP的版本  取值范围：4、6，分别表示创建ipv4和ipv6  约束：必须是系统具体支持的类型  不填或空字符串时，默认创建ipv4
	IpVersion *CreatePublicipOptionIpVersion `json:"ip_version,omitempty"`

	// 功能说明：弹性公网IP名称 取值范围：1-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）
	Alias *string `json:"alias,omitempty"`

	// 功能说明：端口id  约束：必须是存在的端口id，如果该端口不存在或端口已绑定EIP则会提示出错。
	PortId *string `json:"port_id,omitempty"`
}

func (o CreatePublicipOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicipOption struct{}"
	}

	return strings.Join([]string{"CreatePublicipOption", string(data)}, " ")
}

type CreatePublicipOptionIpVersion struct {
	value int32
}

type CreatePublicipOptionIpVersionEnum struct {
	E_4 CreatePublicipOptionIpVersion
	E_6 CreatePublicipOptionIpVersion
}

func GetCreatePublicipOptionIpVersionEnum() CreatePublicipOptionIpVersionEnum {
	return CreatePublicipOptionIpVersionEnum{
		E_4: CreatePublicipOptionIpVersion{
			value: 4,
		}, E_6: CreatePublicipOptionIpVersion{
			value: 6,
		},
	}
}

func (c CreatePublicipOptionIpVersion) Value() int32 {
	return c.value
}

func (c CreatePublicipOptionIpVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePublicipOptionIpVersion) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
