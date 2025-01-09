package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApplyDesktopsInternet 开通桌面上网能力请求体。
type ApplyDesktopsInternet struct {

	// 需要开通上网功能的桌面id列表。
	DesktopIds *[]string `json:"desktop_ids,omitempty"`

	// 支持的类型请参考EIP服务支持的类型。可通过调用如下链接的接口查询，https://support.huaweicloud.com/api-eip/ShowPublicIpType.html。
	EipType string `json:"eip_type"`

	// eip带宽计费模式 - TRAFFIC：按流量计费。 - BANDWIDTH：按带宽计费。
	EipChargeMode ApplyDesktopsInternetEipChargeMode `json:"eip_charge_mode"`

	// 带宽大小，单位Mbit/s。默认1Mbit/s~2000Mbit/s（具体范围以各区域配置为准，请参见控制台对应页面显示）。
	BandwidthSize int32 `json:"bandwidth_size"`

	// 需要购买EIP的数量，当desktop_ids为空时需要填，兼容单独购买EIP场景。
	Count *int32 `json:"count,omitempty"`
}

func (o ApplyDesktopsInternet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyDesktopsInternet struct{}"
	}

	return strings.Join([]string{"ApplyDesktopsInternet", string(data)}, " ")
}

type ApplyDesktopsInternetEipChargeMode struct {
	value string
}

type ApplyDesktopsInternetEipChargeModeEnum struct {
	TRAFFIC   ApplyDesktopsInternetEipChargeMode
	BANDWIDTH ApplyDesktopsInternetEipChargeMode
}

func GetApplyDesktopsInternetEipChargeModeEnum() ApplyDesktopsInternetEipChargeModeEnum {
	return ApplyDesktopsInternetEipChargeModeEnum{
		TRAFFIC: ApplyDesktopsInternetEipChargeMode{
			value: "TRAFFIC",
		},
		BANDWIDTH: ApplyDesktopsInternetEipChargeMode{
			value: "BANDWIDTH",
		},
	}
}

func (c ApplyDesktopsInternetEipChargeMode) Value() string {
	return c.value
}

func (c ApplyDesktopsInternetEipChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApplyDesktopsInternetEipChargeMode) UnmarshalJSON(b []byte) error {
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
