package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdatePremiumInstanceResponse Response Object
type UpdatePremiumInstanceResponse struct {

	// **参数解释：** 独享引擎ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释：** 独享引擎名称 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Instancename *string `json:"instancename,omitempty"`

	// **参数解释：** 独享引擎ECS ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ServerId *string `json:"serverId,omitempty"`

	// **参数解释：** Region代码 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释：** 可用区代码 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Zone *string `json:"zone,omitempty"`

	// **参数解释：** CPU架构代码 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Arch *string `json:"arch,omitempty"`

	// **参数解释：** ECS规格代码 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	CpuFlavor *string `json:"cpu_flavor,omitempty"`

	// **参数解释：** 独享引擎所在VPC ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释：** 独享引擎所在VPC的子网ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	SubnetId *string `json:"subnet_id,omitempty"`

	// **参数解释：** 独享引擎的业务面IP **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ServiceIp *string `json:"service_ip,omitempty"`

	// **参数解释：** 独享引擎的业务面IPV6地址 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ServiceIpv6 *string `json:"service_ipv6,omitempty"`

	// **参数解释：** 独享引擎的管理面IP **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	FloatIp *string `json:"floatIp,omitempty"`

	// **参数解释：** 独享引擎ECS绑定的安全组 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	SecurityGroupIds *[]string `json:"security_group_ids,omitempty"`

	// **参数解释：** 独享引擎计费状态标识，用于指示独享引擎当前的计费使用状态 **约束限制：** 不涉及 **取值范围：**  - 0：正常计费  - 1：冻结，资源和数据会保留，但租户无法再正常使用云服务  - 2：终止，资源和数据将清除 **默认取值：** 不涉及
	Status *UpdatePremiumInstanceResponseStatus `json:"status,omitempty"`

	// **参数解释：** 独享引擎运行状态标识，用于反映独享引擎当前的运行生命周期状态 **约束限制：** 不涉及 **取值范围：**  - 0：创建中  - 1：运行中  - 2：删除中  - 3：已删除  - 4：创建失败  - 5：已冻结  - 6：异常  - 7：更新中  - 8：更新失败 **默认取值：** 不涉及
	RunStatus *UpdatePremiumInstanceResponseRunStatus `json:"run_status,omitempty"`

	// **参数解释：** 独享引擎接入状态 **约束限制：** 不涉及 **取值范围：**  - 0: 未接入  - 1: 已接入  **默认取值：** 不涉及
	AccessStatus *UpdatePremiumInstanceResponseAccessStatus `json:"access_status,omitempty"`

	// **参数解释：** 独享引擎是否可升级（0：不可升级，1：可升级） **约束限制：** 不涉及 **取值范围：**  - 0: 不可升级  - 1: 可升级  **默认取值：** 不涉及
	Upgradable *UpdatePremiumInstanceResponseUpgradable `json:"upgradable,omitempty"`

	// **参数解释：** 云服务代码 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	CloudServiceType *string `json:"cloudServiceType,omitempty"`

	// **参数解释：** 云服务资源类型 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ResourceType *string `json:"resourceType,omitempty"`

	// **参数解释：** 云服务资源代码 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ResourceSpecCode *string `json:"resourceSpecCode,omitempty"`

	// **参数解释：** 独享引擎ECS规格，如\"8vCPUs | 16GB\" **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Specification *string `json:"specification,omitempty"`

	// **参数解释：** 独享引擎防护的域名 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Hosts *[]IdHostnameEntry `json:"hosts,omitempty"`

	// **参数解释：** 存储类型（可选） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	VolumeType *string `json:"volume_type,omitempty"`

	// **参数解释：** 存储资源池ID（可选） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释：** 独享引擎所在WAF组的ID（仅适用特殊独享模式） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PoolId *string `json:"pool_id,omitempty"`

	// **参数解释：** '计费模式。0: 包周期；1:按需' **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ChargeMode     *int32 `json:"charge_mode,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdatePremiumInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePremiumInstanceResponse struct{}"
	}

	return strings.Join([]string{"UpdatePremiumInstanceResponse", string(data)}, " ")
}

type UpdatePremiumInstanceResponseStatus struct {
	value int32
}

type UpdatePremiumInstanceResponseStatusEnum struct {
	E_0 UpdatePremiumInstanceResponseStatus
	E_1 UpdatePremiumInstanceResponseStatus
	E_2 UpdatePremiumInstanceResponseStatus
}

func GetUpdatePremiumInstanceResponseStatusEnum() UpdatePremiumInstanceResponseStatusEnum {
	return UpdatePremiumInstanceResponseStatusEnum{
		E_0: UpdatePremiumInstanceResponseStatus{
			value: 0,
		}, E_1: UpdatePremiumInstanceResponseStatus{
			value: 1,
		}, E_2: UpdatePremiumInstanceResponseStatus{
			value: 2,
		},
	}
}

func (c UpdatePremiumInstanceResponseStatus) Value() int32 {
	return c.value
}

func (c UpdatePremiumInstanceResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePremiumInstanceResponseStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type UpdatePremiumInstanceResponseRunStatus struct {
	value int32
}

type UpdatePremiumInstanceResponseRunStatusEnum struct {
	E_0 UpdatePremiumInstanceResponseRunStatus
	E_1 UpdatePremiumInstanceResponseRunStatus
	E_2 UpdatePremiumInstanceResponseRunStatus
	E_3 UpdatePremiumInstanceResponseRunStatus
	E_4 UpdatePremiumInstanceResponseRunStatus
	E_5 UpdatePremiumInstanceResponseRunStatus
	E_6 UpdatePremiumInstanceResponseRunStatus
	E_7 UpdatePremiumInstanceResponseRunStatus
	E_8 UpdatePremiumInstanceResponseRunStatus
}

func GetUpdatePremiumInstanceResponseRunStatusEnum() UpdatePremiumInstanceResponseRunStatusEnum {
	return UpdatePremiumInstanceResponseRunStatusEnum{
		E_0: UpdatePremiumInstanceResponseRunStatus{
			value: 0,
		}, E_1: UpdatePremiumInstanceResponseRunStatus{
			value: 1,
		}, E_2: UpdatePremiumInstanceResponseRunStatus{
			value: 2,
		}, E_3: UpdatePremiumInstanceResponseRunStatus{
			value: 3,
		}, E_4: UpdatePremiumInstanceResponseRunStatus{
			value: 4,
		}, E_5: UpdatePremiumInstanceResponseRunStatus{
			value: 5,
		}, E_6: UpdatePremiumInstanceResponseRunStatus{
			value: 6,
		}, E_7: UpdatePremiumInstanceResponseRunStatus{
			value: 7,
		}, E_8: UpdatePremiumInstanceResponseRunStatus{
			value: 8,
		},
	}
}

func (c UpdatePremiumInstanceResponseRunStatus) Value() int32 {
	return c.value
}

func (c UpdatePremiumInstanceResponseRunStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePremiumInstanceResponseRunStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type UpdatePremiumInstanceResponseAccessStatus struct {
	value int32
}

type UpdatePremiumInstanceResponseAccessStatusEnum struct {
	E_0 UpdatePremiumInstanceResponseAccessStatus
	E_1 UpdatePremiumInstanceResponseAccessStatus
}

func GetUpdatePremiumInstanceResponseAccessStatusEnum() UpdatePremiumInstanceResponseAccessStatusEnum {
	return UpdatePremiumInstanceResponseAccessStatusEnum{
		E_0: UpdatePremiumInstanceResponseAccessStatus{
			value: 0,
		}, E_1: UpdatePremiumInstanceResponseAccessStatus{
			value: 1,
		},
	}
}

func (c UpdatePremiumInstanceResponseAccessStatus) Value() int32 {
	return c.value
}

func (c UpdatePremiumInstanceResponseAccessStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePremiumInstanceResponseAccessStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type UpdatePremiumInstanceResponseUpgradable struct {
	value int32
}

type UpdatePremiumInstanceResponseUpgradableEnum struct {
	E_0 UpdatePremiumInstanceResponseUpgradable
	E_1 UpdatePremiumInstanceResponseUpgradable
}

func GetUpdatePremiumInstanceResponseUpgradableEnum() UpdatePremiumInstanceResponseUpgradableEnum {
	return UpdatePremiumInstanceResponseUpgradableEnum{
		E_0: UpdatePremiumInstanceResponseUpgradable{
			value: 0,
		}, E_1: UpdatePremiumInstanceResponseUpgradable{
			value: 1,
		},
	}
}

func (c UpdatePremiumInstanceResponseUpgradable) Value() int32 {
	return c.value
}

func (c UpdatePremiumInstanceResponseUpgradable) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePremiumInstanceResponseUpgradable) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
