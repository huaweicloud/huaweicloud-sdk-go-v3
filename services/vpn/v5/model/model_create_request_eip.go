package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateRequestEip struct {

	// 功能说明：公网IP的唯一标识
	Id *string `json:"id,omitempty"`

	// 功能说明：EIP的类型  取值范围：5_bgp（全动态BGP），5_sbgp（静态BGP）  华南-广州：5_bgp、5_sbgp  华东-上海一：5_bgp、5_sbgp  华东-上海二：5_bgp、5_sbgp  华北-北京一：5_bgp、5_sbgp  中国-香港：5_bgp  亚太-曼谷：5_bgp  亚太-新加坡：5_bgp  非洲-约翰内斯堡：5_bgp  西南-贵阳一：5_bgp、5_sbgp  华北-北京四：5_bgp、5_sbgp  拉美-圣地亚哥：5_bgp  拉美-圣保罗一：5_bgp  拉美-墨西哥城一：5_bgp  拉美-布宜诺斯艾利一：5_bgp  拉美-利马一：5_bgp  拉美-圣地亚哥二： 5_bgp 约束：必须是系统具体支持的类型。
	Type *string `json:"type,omitempty"`

	// 功能说明：按流量计费还是按带宽计费  取值范围：  bandwidth：按带宽计费  traffic：按流量计费  95peak_plus：按增强型95计费
	ChargeMode *CreateRequestEipChargeMode `json:"charge_mode,omitempty"`

	// 带宽大小Mbit/s。flavor为Basic时，取值不能大于100；flavor为Professional1时，取值不能大于300；flavor为Professional2时，取值不能大于1024
	BandwidthSize *int32 `json:"bandwidth_size,omitempty"`

	// 带宽名称
	BandwidthName *string `json:"bandwidth_name,omitempty"`
}

func (o CreateRequestEip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRequestEip struct{}"
	}

	return strings.Join([]string{"CreateRequestEip", string(data)}, " ")
}

type CreateRequestEipChargeMode struct {
	value string
}

type CreateRequestEipChargeModeEnum struct {
	BANDWIDTH CreateRequestEipChargeMode
	TRAFFIC   CreateRequestEipChargeMode
}

func GetCreateRequestEipChargeModeEnum() CreateRequestEipChargeModeEnum {
	return CreateRequestEipChargeModeEnum{
		BANDWIDTH: CreateRequestEipChargeMode{
			value: "bandwidth",
		},
		TRAFFIC: CreateRequestEipChargeMode{
			value: "traffic",
		},
	}
}

func (c CreateRequestEipChargeMode) Value() string {
	return c.value
}

func (c CreateRequestEipChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRequestEipChargeMode) UnmarshalJSON(b []byte) error {
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
