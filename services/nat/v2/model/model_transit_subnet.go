package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// TransitSubnet 中转子网的响应体
type TransitSubnet struct {

	// 中转子网的ID
	Id string `json:"id"`

	// 中转子网的名称
	Name string `json:"name"`

	// 中转子网的描述
	Description string `json:"description"`

	// 中转子网的子网所属的项目ID
	VirsubnetProjectId string `json:"virsubnet_project_id"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 中转子网的子网所属VPC的ID
	VpcId string `json:"vpc_id"`

	// 中转子网的子网ID
	VirsubnetId string `json:"virsubnet_id"`

	// 中转子网的子网网段
	Cidr string `json:"cidr"`

	// 中转子网类型。取值范围：VPC
	Type TransitSubnetType `json:"type"`

	// 中转子网状态。 取值范围： ACTIVE： 当前资源状态正常。 INACTIVE： 不可用。
	Status TransitSubnetStatus `json:"status"`

	// 当前中转子网下已分配的中转子网IP数量。
	IpCount int32 `json:"ip_count"`

	// 中转子网创建时间，遵循UTC时间，格式是yyyy-mm-ddThh:mm:ssZ。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 中转子网更新时间，遵循UTC时间，格式是yyyy-mm-ddThh:mm:ssZ。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 标签列表
	Tags []Tag `json:"tags"`
}

func (o TransitSubnet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransitSubnet struct{}"
	}

	return strings.Join([]string{"TransitSubnet", string(data)}, " ")
}

type TransitSubnetType struct {
	value string
}

type TransitSubnetTypeEnum struct {
	VPC TransitSubnetType
}

func GetTransitSubnetTypeEnum() TransitSubnetTypeEnum {
	return TransitSubnetTypeEnum{
		VPC: TransitSubnetType{
			value: "VPC",
		},
	}
}

func (c TransitSubnetType) Value() string {
	return c.value
}

func (c TransitSubnetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TransitSubnetType) UnmarshalJSON(b []byte) error {
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

type TransitSubnetStatus struct {
	value string
}

type TransitSubnetStatusEnum struct {
	ACTIVE   TransitSubnetStatus
	INACTIVE TransitSubnetStatus
}

func GetTransitSubnetStatusEnum() TransitSubnetStatusEnum {
	return TransitSubnetStatusEnum{
		ACTIVE: TransitSubnetStatus{
			value: "ACTIVE",
		},
		INACTIVE: TransitSubnetStatus{
			value: "INACTIVE",
		},
	}
}

func (c TransitSubnetStatus) Value() string {
	return c.value
}

func (c TransitSubnetStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TransitSubnetStatus) UnmarshalJSON(b []byte) error {
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
