package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// HostedCloud 归属云。
type HostedCloud struct {

	// - HWCloud (华为云) - Ireland (爱尔兰)
	HostedCloud HostedCloudHostedCloud `json:"hosted_cloud"`
}

func (o HostedCloud) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostedCloud struct{}"
	}

	return strings.Join([]string{"HostedCloud", string(data)}, " ")
}

type HostedCloudHostedCloud struct {
	value string
}

type HostedCloudHostedCloudEnum struct {
	HW_CLOUD HostedCloudHostedCloud
	IRELAND  HostedCloudHostedCloud
}

func GetHostedCloudHostedCloudEnum() HostedCloudHostedCloudEnum {
	return HostedCloudHostedCloudEnum{
		HW_CLOUD: HostedCloudHostedCloud{
			value: "HWCloud",
		},
		IRELAND: HostedCloudHostedCloud{
			value: "Ireland",
		},
	}
}

func (c HostedCloudHostedCloud) Value() string {
	return c.value
}

func (c HostedCloudHostedCloud) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HostedCloudHostedCloud) UnmarshalJSON(b []byte) error {
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
