package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type OpenEngressEipReq struct {

	// 出公网带宽  单位：Mbit/s
	BandwidthSize *string `json:"bandwidth_size,omitempty"`

	// 出公网带宽计费类型： - bandwidth：按带宽计费 - traffic：按流量计费
	BandwidthChargingMode *OpenEngressEipReqBandwidthChargingMode `json:"bandwidth_charging_mode,omitempty"`
}

func (o OpenEngressEipReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenEngressEipReq struct{}"
	}

	return strings.Join([]string{"OpenEngressEipReq", string(data)}, " ")
}

type OpenEngressEipReqBandwidthChargingMode struct {
	value string
}

type OpenEngressEipReqBandwidthChargingModeEnum struct {
	BANDWIDTH OpenEngressEipReqBandwidthChargingMode
	TRAFFIC   OpenEngressEipReqBandwidthChargingMode
}

func GetOpenEngressEipReqBandwidthChargingModeEnum() OpenEngressEipReqBandwidthChargingModeEnum {
	return OpenEngressEipReqBandwidthChargingModeEnum{
		BANDWIDTH: OpenEngressEipReqBandwidthChargingMode{
			value: "bandwidth",
		},
		TRAFFIC: OpenEngressEipReqBandwidthChargingMode{
			value: "traffic",
		},
	}
}

func (c OpenEngressEipReqBandwidthChargingMode) Value() string {
	return c.value
}

func (c OpenEngressEipReqBandwidthChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenEngressEipReqBandwidthChargingMode) UnmarshalJSON(b []byte) error {
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
