package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EipResult eip信息
type EipResult struct {

	// IP地址类型。
	IpType *EipResultIpType `json:"ip_type,omitempty"`

	Bandwidth *BandwidthResult `json:"bandwidth,omitempty"`
}

func (o EipResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipResult struct{}"
	}

	return strings.Join([]string{"EipResult", string(data)}, " ")
}

type EipResultIpType struct {
	value string
}

type EipResultIpTypeEnum struct {
	E_5_BGP    EipResultIpType
	E_5_SBGP   EipResultIpType
	E_5_TELCOM EipResultIpType
	E_5_UNION  EipResultIpType
}

func GetEipResultIpTypeEnum() EipResultIpTypeEnum {
	return EipResultIpTypeEnum{
		E_5_BGP: EipResultIpType{
			value: "5_bgp",
		},
		E_5_SBGP: EipResultIpType{
			value: "5_sbgp",
		},
		E_5_TELCOM: EipResultIpType{
			value: "5_telcom",
		},
		E_5_UNION: EipResultIpType{
			value: "5_union",
		},
	}
}

func (c EipResultIpType) Value() string {
	return c.value
}

func (c EipResultIpType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EipResultIpType) UnmarshalJSON(b []byte) error {
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
