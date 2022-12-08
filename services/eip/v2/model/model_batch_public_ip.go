package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// publicip信息
type BatchPublicIp struct {

	// 指定id创建EIP
	Id *string `json:"id,omitempty"`

	// 公网ip类型
	Type BatchPublicIpType `json:"type"`

	// 公网EIP的版本，例如ipv4，ipv6，默认为ipv4
	IpVersion *string `json:"ip_version,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 公网EIP标签
	Tags *[]string `json:"tags,omitempty"`

	Profile *BatchProfile `json:"profile,omitempty"`
}

func (o BatchPublicIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchPublicIp struct{}"
	}

	return strings.Join([]string{"BatchPublicIp", string(data)}, " ")
}

type BatchPublicIpType struct {
	value string
}

type BatchPublicIpTypeEnum struct {
	E_5_BGP   BatchPublicIpType
	E_5_UNION BatchPublicIpType
	E_5_SBGP  BatchPublicIpType
}

func GetBatchPublicIpTypeEnum() BatchPublicIpTypeEnum {
	return BatchPublicIpTypeEnum{
		E_5_BGP: BatchPublicIpType{
			value: "5_bgp",
		},
		E_5_UNION: BatchPublicIpType{
			value: "5_union",
		},
		E_5_SBGP: BatchPublicIpType{
			value: "5_sbgp",
		},
	}
}

func (c BatchPublicIpType) Value() string {
	return c.value
}

func (c BatchPublicIpType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchPublicIpType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
