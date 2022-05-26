package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// 云连接实例。
type CloudConnection struct {

	// 云连接实例的ID。
	Id *string `json:"id,omitempty"`

	// 云连接实例的名字。
	Name *string `json:"name,omitempty"`

	// 云连接实例的描述。
	Description *string `json:"description,omitempty"`

	// 帐号ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 云连接实例的企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 云连接实例的状态。ACTIVE：表示状态可用。
	Status *CloudConnectionStatus `json:"status,omitempty"`

	// 云连接实例的管理状态。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 云连接实例的创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 云连接实例的更新时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 云连接使用场景。|- VPC：虚拟私有云。 ER：虚拟路由器。
	UsedScene *CloudConnectionUsedScene `json:"used_scene,omitempty"`

	// 云连接实例关联网络实例的个数。
	NetworkInstanceNumber *int32 `json:"network_instance_number,omitempty"`

	// 云连接实例关联带宽包的个数。
	BandwidthPackageNumber *int32 `json:"bandwidth_package_number,omitempty"`

	// 云连接实例关联域间带宽的个数。
	InterRegionBandwidthNumber *int32 `json:"inter_region_bandwidth_number,omitempty"`
}

func (o CloudConnection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudConnection struct{}"
	}

	return strings.Join([]string{"CloudConnection", string(data)}, " ")
}

type CloudConnectionStatus struct {
	value string
}

type CloudConnectionStatusEnum struct {
	ACTIVE CloudConnectionStatus
}

func GetCloudConnectionStatusEnum() CloudConnectionStatusEnum {
	return CloudConnectionStatusEnum{
		ACTIVE: CloudConnectionStatus{
			value: "ACTIVE",
		},
	}
}

func (c CloudConnectionStatus) Value() string {
	return c.value
}

func (c CloudConnectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CloudConnectionStatus) UnmarshalJSON(b []byte) error {
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

type CloudConnectionUsedScene struct {
	value string
}

type CloudConnectionUsedSceneEnum struct {
	ER  CloudConnectionUsedScene
	VPC CloudConnectionUsedScene
}

func GetCloudConnectionUsedSceneEnum() CloudConnectionUsedSceneEnum {
	return CloudConnectionUsedSceneEnum{
		ER: CloudConnectionUsedScene{
			value: "er",
		},
		VPC: CloudConnectionUsedScene{
			value: "vpc",
		},
	}
}

func (c CloudConnectionUsedScene) Value() string {
	return c.value
}

func (c CloudConnectionUsedScene) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CloudConnectionUsedScene) UnmarshalJSON(b []byte) error {
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
