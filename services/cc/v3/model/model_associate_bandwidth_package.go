package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AssociateBandwidthPackage 将带宽包实例关联到资源的详细信息。
type AssociateBandwidthPackage struct {

	// 带宽包实例绑定的资源ID。
	ResourceId string `json:"resource_id"`

	// 带宽包实例绑定的资源类型。 cloud_connection: 云连接实例。
	ResourceType AssociateBandwidthPackageResourceType `json:"resource_type"`
}

func (o AssociateBandwidthPackage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateBandwidthPackage struct{}"
	}

	return strings.Join([]string{"AssociateBandwidthPackage", string(data)}, " ")
}

type AssociateBandwidthPackageResourceType struct {
	value string
}

type AssociateBandwidthPackageResourceTypeEnum struct {
	CLOUD_CONNECTION AssociateBandwidthPackageResourceType
}

func GetAssociateBandwidthPackageResourceTypeEnum() AssociateBandwidthPackageResourceTypeEnum {
	return AssociateBandwidthPackageResourceTypeEnum{
		CLOUD_CONNECTION: AssociateBandwidthPackageResourceType{
			value: "cloud_connection",
		},
	}
}

func (c AssociateBandwidthPackageResourceType) Value() string {
	return c.value
}

func (c AssociateBandwidthPackageResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AssociateBandwidthPackageResourceType) UnmarshalJSON(b []byte) error {
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
