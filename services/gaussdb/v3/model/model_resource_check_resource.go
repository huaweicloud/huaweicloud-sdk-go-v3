package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResourceCheckResource 资源信息。
type ResourceCheckResource struct {

	// 企业项目ID。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 实例数量。
	InstanceNum int32 `json:"instance_num"`

	// 实例部署模式。
	Mode ResourceCheckResourceMode `json:"mode"`

	// 可用区类型，目前仅支持single。
	AvailabilityZoneMode string `json:"availability_zone_mode"`

	// FE节点数量。
	FeNodeNum int32 `json:"fe_node_num"`

	// BE节点数量。
	BeNodeNum int32 `json:"be_node_num"`

	// FE规格码。
	FeFlavorRef string `json:"fe_flavor_ref"`

	// BE规格码。
	BeFlavorRef string `json:"be_flavor_ref"`

	// 可用区码。选填，校验可用区码是否正确。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// HTAP实例子网即GaussDBForMySQL实例子网。 获取方法请参见[获取子网ID](https://support.huaweicloud.com/api-vpc/vpc_subnet01_0003.html)。
	SubnetId string `json:"subnet_id"`
}

func (o ResourceCheckResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceCheckResource struct{}"
	}

	return strings.Join([]string{"ResourceCheckResource", string(data)}, " ")
}

type ResourceCheckResourceMode struct {
	value string
}

type ResourceCheckResourceModeEnum struct {
	CLUSTER ResourceCheckResourceMode
	SINGLE  ResourceCheckResourceMode
}

func GetResourceCheckResourceModeEnum() ResourceCheckResourceModeEnum {
	return ResourceCheckResourceModeEnum{
		CLUSTER: ResourceCheckResourceMode{
			value: "Cluster",
		},
		SINGLE: ResourceCheckResourceMode{
			value: "Single",
		},
	}
}

func (c ResourceCheckResourceMode) Value() string {
	return c.value
}

func (c ResourceCheckResourceMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceCheckResourceMode) UnmarshalJSON(b []byte) error {
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
