package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type OpenIngressEipReq struct {

	// 入公网带宽  单位：Mbit/s
	BandwidthSize *int32 `json:"bandwidth_size,omitempty"`

	// 入公网带宽计费类型： - bandwidth：按带宽计费 - [traffic：按流量计费](tag:hws_test)
	BandwidthChargingMode *OpenIngressEipReqBandwidthChargingMode `json:"bandwidth_charging_mode,omitempty"`
}

func (o OpenIngressEipReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenIngressEipReq struct{}"
	}

	return strings.Join([]string{"OpenIngressEipReq", string(data)}, " ")
}

type OpenIngressEipReqBandwidthChargingMode struct {
	value string
}

type OpenIngressEipReqBandwidthChargingModeEnum struct {
	BANDWIDTH OpenIngressEipReqBandwidthChargingMode
	TRAFFIC   OpenIngressEipReqBandwidthChargingMode
}

func GetOpenIngressEipReqBandwidthChargingModeEnum() OpenIngressEipReqBandwidthChargingModeEnum {
	return OpenIngressEipReqBandwidthChargingModeEnum{
		BANDWIDTH: OpenIngressEipReqBandwidthChargingMode{
			value: "bandwidth",
		},
		TRAFFIC: OpenIngressEipReqBandwidthChargingMode{
			value: "traffic",
		},
	}
}

func (c OpenIngressEipReqBandwidthChargingMode) Value() string {
	return c.value
}

func (c OpenIngressEipReqBandwidthChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenIngressEipReqBandwidthChargingMode) UnmarshalJSON(b []byte) error {
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
