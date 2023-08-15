package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PublicipInfoResponseBody PublicipInfo对象
type PublicipInfoResponseBody struct {

	// - 功能说明：弹性公网IPV4或IPv6的公网地址
	PublicipAddress *string `json:"publicip_address,omitempty"`

	// - 功能说明：带宽对应的弹性公网IPV4或IPv6的唯一标识
	PublicipId *string `json:"publicip_id,omitempty"`

	// - 功能说明：EIP的类型  - [取值范围：5_bgp](tag:hk_g42,g42,hk_sbc,sbc)  - [取值范围：5_chinamobile](tag:cmcc)  - [取值范围：5_bgp（Dynamic BGP）、5_mainbgp（Mail BGP）](tag:dt)  - [取值范围：    - eu-de: 5_bgp（Dynamic BGP）、5_mainbgp（Mail BGP）、5_gray（Dedicated Load Balancer）    - eu-nl: 5_bgp（Dynamic BGP）、5_mainbgp（Mail BGP）](tag:dt_test) - [取值范围：5_bgp（全动态BGP），5_sbgp（静态BGP），5_youxuanbgp（优选BGP）    - 华南-广州：5_bgp、5_sbgp    - 华东-上海一：5_bgp、5_sbgp    - 华东-上海二：5_bgp、5_sbgp    - 华北-北京一：5_bgp、5_sbgp    - 中国-香港：5_bgp、5_youxuanbgp    - 亚太-曼谷：5_bgp    - 亚太-新加坡：5_bgp    - 非洲-约翰内斯堡：5_bgp    - 西南-贵阳一：5_sbgp    - 华北-北京四：5_bgp、5_sbgp    - 拉美-圣地亚哥：5_bgp    - 拉美-圣保罗一：5_bgp    - 拉美-墨西哥城一：5_bgp    - 拉美-布宜诺斯艾利一：5_bgp    - 拉美-利马一：5_bgp    - 拉美-圣地亚哥二： 5_bgp  ](tag:hws,hws_hk)  - 约束：     - 必须是系统具体支持的类型。     - publicip_id为IPv4端口，所以[\"publicip_type\"](tag:cmcc,ctc,dt,dt_test,fcs,fcs_dt,fcs_vm,hws_hk,hws_ocb,ocb,tlf,tm,hk_g42,g42,hk_sbc,sbc)[\"type\"](tag:hws)字段未给定时，默认为[5_bgp](tag:ctc,dt,dt_test,fcs,fcs_dt,fcs_vm,g42,hk_g42,hk_sbc,hws,hws_hk,hws_ocb,mix,ocb,sbc,tlf,tm)[5_chinamobile](tag:cmcc)。
	PublicipType *string `json:"publicip_type,omitempty"`

	// - 功能说明：IPv4时无此字段，IPv6时为申请到的弹性公网IP地址
	Publicipv6Address *string `json:"publicipv6_address,omitempty"`

	// - 功能说明：IP版本信息 - 取值范围：  4：IPv4；  6：IPv6
	IpVersion *PublicipInfoResponseBodyIpVersion `json:"ip_version,omitempty"`
}

func (o PublicipInfoResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicipInfoResponseBody struct{}"
	}

	return strings.Join([]string{"PublicipInfoResponseBody", string(data)}, " ")
}

type PublicipInfoResponseBodyIpVersion struct {
	value int32
}

type PublicipInfoResponseBodyIpVersionEnum struct {
	E_4 PublicipInfoResponseBodyIpVersion
	E_6 PublicipInfoResponseBodyIpVersion
}

func GetPublicipInfoResponseBodyIpVersionEnum() PublicipInfoResponseBodyIpVersionEnum {
	return PublicipInfoResponseBodyIpVersionEnum{
		E_4: PublicipInfoResponseBodyIpVersion{
			value: 4,
		}, E_6: PublicipInfoResponseBodyIpVersion{
			value: 6,
		},
	}
}

func (c PublicipInfoResponseBodyIpVersion) Value() int32 {
	return c.value
}

func (c PublicipInfoResponseBodyIpVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicipInfoResponseBodyIpVersion) UnmarshalJSON(b []byte) error {
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
