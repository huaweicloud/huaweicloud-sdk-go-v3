package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowInstanceResponse Response Object
type ShowInstanceResponse struct {

	// 实例ID
	Id *string `json:"id,omitempty"`

	// 实例名称
	Name *string `json:"name,omitempty"`

	// 实例描述
	Description *string `json:"description,omitempty"`

	// 用户虚拟私有云ID
	VpcId *string `json:"vpc_id,omitempty"`

	// 用户子网的网络ID
	SubnetId *string `json:"subnet_id,omitempty"`

	// 实例规格
	Spec *string `json:"spec,omitempty"`

	// 实例版本
	Version *string `json:"version,omitempty"`

	// 实例计费模式
	ChargeMode *ShowInstanceResponseChargeMode `json:"charge_mode,omitempty"`

	// 访问地址
	AccessAddress *string `json:"access_address,omitempty"`

	// 实例创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 实例更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 实例过期时间
	ExpiresAt *string `json:"expires_at,omitempty"`

	// 实例状态 取值范围： - Initial：初始化 - Creating：创建中 - Running：运行中 - Unavailable：不可用 - Frozen：已冻结 - DeploymentAbnormal：部署异常 - PartiallyUnavailable：部分不可用 - CoreDialFailed：核心拨测失败
	Status *ShowInstanceResponseStatus `json:"status,omitempty"`

	// obs桶名称
	ObsBucketName *string `json:"obs_bucket_name,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 是否为用户指定obs桶
	UserDefObs *bool `json:"user_def_obs,omitempty"`

	// 产品ID
	ProductId *string `json:"product_id,omitempty"`

	// 订单ID，包周期预留字段
	OrderId *string `json:"order_id,omitempty"`

	// VPC的名称
	VpcName *string `json:"vpc_name,omitempty"`

	// VPC的可用子网的范围
	VpcCidr *string `json:"vpc_cidr,omitempty"`

	// VPC的子网名称
	SubnetName *string `json:"subnet_name,omitempty"`

	// 子网的网段
	SubnetCidr *string `json:"subnet_cidr,omitempty"`

	// 实例对应的VPC终端节点服务ID
	VpcepServiceId *string `json:"vpcep_service_id,omitempty"`

	// 标签列表
	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceResponse", string(data)}, " ")
}

type ShowInstanceResponseChargeMode struct {
	value string
}

type ShowInstanceResponseChargeModeEnum struct {
	POST_PAID ShowInstanceResponseChargeMode
}

func GetShowInstanceResponseChargeModeEnum() ShowInstanceResponseChargeModeEnum {
	return ShowInstanceResponseChargeModeEnum{
		POST_PAID: ShowInstanceResponseChargeMode{
			value: "postPaid",
		},
	}
}

func (c ShowInstanceResponseChargeMode) Value() string {
	return c.value
}

func (c ShowInstanceResponseChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceResponseChargeMode) UnmarshalJSON(b []byte) error {
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

type ShowInstanceResponseStatus struct {
	value string
}

type ShowInstanceResponseStatusEnum struct {
	INITIAL               ShowInstanceResponseStatus
	CREATING              ShowInstanceResponseStatus
	RUNNING               ShowInstanceResponseStatus
	UNAVAILABLE           ShowInstanceResponseStatus
	FROZEN                ShowInstanceResponseStatus
	DEPLOYMENT_ABNORMAL   ShowInstanceResponseStatus
	PARTIALLY_UNAVAILABLE ShowInstanceResponseStatus
	CORE_DIAL_FAILED      ShowInstanceResponseStatus
}

func GetShowInstanceResponseStatusEnum() ShowInstanceResponseStatusEnum {
	return ShowInstanceResponseStatusEnum{
		INITIAL: ShowInstanceResponseStatus{
			value: "Initial",
		},
		CREATING: ShowInstanceResponseStatus{
			value: "Creating",
		},
		RUNNING: ShowInstanceResponseStatus{
			value: "Running",
		},
		UNAVAILABLE: ShowInstanceResponseStatus{
			value: "Unavailable",
		},
		FROZEN: ShowInstanceResponseStatus{
			value: "Frozen",
		},
		DEPLOYMENT_ABNORMAL: ShowInstanceResponseStatus{
			value: "DeploymentAbnormal",
		},
		PARTIALLY_UNAVAILABLE: ShowInstanceResponseStatus{
			value: "PartiallyUnavailable",
		},
		CORE_DIAL_FAILED: ShowInstanceResponseStatus{
			value: "CoreDialFailed",
		},
	}
}

func (c ShowInstanceResponseStatus) Value() string {
	return c.value
}

func (c ShowInstanceResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceResponseStatus) UnmarshalJSON(b []byte) error {
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
