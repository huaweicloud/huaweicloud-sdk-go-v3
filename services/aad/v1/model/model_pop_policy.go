package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PopPolicy 策略阈值详情
type PopPolicy struct {

	// 位置封禁列表
	BlockLocation []string `json:"block_location"`

	// 协议封禁列表
	BlockProtocol []string `json:"block_protocol"`

	BwList *Bw `json:"bw_list"`

	// 是否开启连接防护
	ConnectionProtection bool `json:"connection_protection"`

	// 连接防护列表
	ConnectionProtectionList []string `json:"connection_protection_list"`

	// 指纹数
	FingerprintCount int32 `json:"fingerprint_count"`

	// 端口封禁数
	PortBlockCount int32 `json:"port_block_count"`

	// 水印数
	WatermarkCount int32 `json:"watermark_count"`

	// 是否存在流量
	IfExistTraffic bool `json:"if_exist_traffic"`

	// 固定值ALL
	Pop PopPolicyPop `json:"pop"`
}

func (o PopPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PopPolicy struct{}"
	}

	return strings.Join([]string{"PopPolicy", string(data)}, " ")
}

type PopPolicyPop struct {
	value string
}

type PopPolicyPopEnum struct {
	ALL PopPolicyPop
}

func GetPopPolicyPopEnum() PopPolicyPopEnum {
	return PopPolicyPopEnum{
		ALL: PopPolicyPop{
			value: "ALL",
		},
	}
}

func (c PopPolicyPop) Value() string {
	return c.value
}

func (c PopPolicyPop) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PopPolicyPop) UnmarshalJSON(b []byte) error {
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
