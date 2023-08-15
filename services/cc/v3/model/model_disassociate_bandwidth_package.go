package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DisassociateBandwidthPackage 解关联带宽包实例的详细信息。
type DisassociateBandwidthPackage struct {

	// 带宽包实例待解关联的资源实例ID。
	ResourceId string `json:"resource_id"`

	// 带宽包实例待解关联的资源实例类型，cloud_connection：表示为云连接实例。
	ResourceType DisassociateBandwidthPackageResourceType `json:"resource_type"`
}

func (o DisassociateBandwidthPackage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateBandwidthPackage struct{}"
	}

	return strings.Join([]string{"DisassociateBandwidthPackage", string(data)}, " ")
}

type DisassociateBandwidthPackageResourceType struct {
	value string
}

type DisassociateBandwidthPackageResourceTypeEnum struct {
	CLOUD_CONNECTION DisassociateBandwidthPackageResourceType
}

func GetDisassociateBandwidthPackageResourceTypeEnum() DisassociateBandwidthPackageResourceTypeEnum {
	return DisassociateBandwidthPackageResourceTypeEnum{
		CLOUD_CONNECTION: DisassociateBandwidthPackageResourceType{
			value: "cloud_connection",
		},
	}
}

func (c DisassociateBandwidthPackageResourceType) Value() string {
	return c.value
}

func (c DisassociateBandwidthPackageResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DisassociateBandwidthPackageResourceType) UnmarshalJSON(b []byte) error {
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
