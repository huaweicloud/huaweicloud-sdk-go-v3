package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteInstanceResponse Response Object
type DeleteInstanceResponse struct {

	// 独享引擎实例ID
	Id *string `json:"id,omitempty"`

	// 独享引擎实例名称
	Instancename *string `json:"instancename,omitempty"`

	// 独享引擎实例Region ID
	Region *string `json:"region,omitempty"`

	// 可用区ID
	Zone *string `json:"zone,omitempty"`

	// CPU架构
	Arch *string `json:"arch,omitempty"`

	// ECS规格
	CpuFlavor *string `json:"cpu_flavor,omitempty"`

	// 独享引擎实例所在VPC ID
	VpcId *string `json:"vpc_id,omitempty"`

	// 独享引擎实例所在VPC的子网ID
	SubnetId *string `json:"subnet_id,omitempty"`

	// 独享引擎实例的业务面IP
	ServiceIp *string `json:"service_ip,omitempty"`

	// 独享引擎绑定的安全组
	SecurityGroupIds *[]string `json:"security_group_ids,omitempty"`

	// **参数解释：** 独享引擎计费状态标识，用于指示独享引擎当前的计费使用状态 **约束限制：** 不涉及 **取值范围：**  - 0：正常计费  - 1：冻结，资源和数据会保留，但租户无法再正常使用云服务  - 2：终止，资源和数据将清除 **默认取值：** 不涉及
	Status *DeleteInstanceResponseStatus `json:"status,omitempty"`

	// **参数解释：** 独享引擎运行状态标识，用于反映独享引擎当前的运行生命周期状态 **约束限制：** 不涉及 **取值范围：**  - 0：创建中  - 1：运行中  - 2：删除中  - 3：已删除  - 4：创建失败  - 5：已冻结  - 6：异常  - 7：更新中  - 8：更新失败 **默认取值：** 不涉及
	RunStatus *DeleteInstanceResponseRunStatus `json:"run_status,omitempty"`

	// **参数解释：** 独享引擎接入状态 **约束限制：** 不涉及 **取值范围：**  - 0: 未接入  - 1: 已接入  **默认取值：** 不涉及
	AccessStatus *DeleteInstanceResponseAccessStatus `json:"access_status,omitempty"`

	// 独享引擎是否可升级（0：不可升级，1：可升级）
	Upgradable *int32 `json:"upgradable,omitempty"`

	// 云服务代码。 仅作为标记，用户可忽略。
	CloudServiceType *string `json:"cloudServiceType,omitempty"`

	// 云服务资源类型，仅作为标记，用户可忽略。
	ResourceType *string `json:"resourceType,omitempty"`

	// 云服务资源代码。仅作为标记，用户可忽略。
	ResourceSpecCode *string `json:"resourceSpecCode,omitempty"`

	// 独享引擎ECS规格，如\"8vCPUs | 16GB\"
	Specification *string `json:"specification,omitempty"`

	// 独享引擎ECS ID
	ServerId *string `json:"serverId,omitempty"`

	// 引擎实例创建时间
	CreateTime     *int64 `json:"create_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DeleteInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstanceResponse", string(data)}, " ")
}

type DeleteInstanceResponseStatus struct {
	value int32
}

type DeleteInstanceResponseStatusEnum struct {
	E_0 DeleteInstanceResponseStatus
	E_1 DeleteInstanceResponseStatus
	E_2 DeleteInstanceResponseStatus
}

func GetDeleteInstanceResponseStatusEnum() DeleteInstanceResponseStatusEnum {
	return DeleteInstanceResponseStatusEnum{
		E_0: DeleteInstanceResponseStatus{
			value: 0,
		}, E_1: DeleteInstanceResponseStatus{
			value: 1,
		}, E_2: DeleteInstanceResponseStatus{
			value: 2,
		},
	}
}

func (c DeleteInstanceResponseStatus) Value() int32 {
	return c.value
}

func (c DeleteInstanceResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteInstanceResponseStatus) UnmarshalJSON(b []byte) error {
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

type DeleteInstanceResponseRunStatus struct {
	value int32
}

type DeleteInstanceResponseRunStatusEnum struct {
	E_0 DeleteInstanceResponseRunStatus
	E_1 DeleteInstanceResponseRunStatus
	E_2 DeleteInstanceResponseRunStatus
	E_3 DeleteInstanceResponseRunStatus
	E_4 DeleteInstanceResponseRunStatus
	E_5 DeleteInstanceResponseRunStatus
	E_6 DeleteInstanceResponseRunStatus
	E_7 DeleteInstanceResponseRunStatus
	E_8 DeleteInstanceResponseRunStatus
}

func GetDeleteInstanceResponseRunStatusEnum() DeleteInstanceResponseRunStatusEnum {
	return DeleteInstanceResponseRunStatusEnum{
		E_0: DeleteInstanceResponseRunStatus{
			value: 0,
		}, E_1: DeleteInstanceResponseRunStatus{
			value: 1,
		}, E_2: DeleteInstanceResponseRunStatus{
			value: 2,
		}, E_3: DeleteInstanceResponseRunStatus{
			value: 3,
		}, E_4: DeleteInstanceResponseRunStatus{
			value: 4,
		}, E_5: DeleteInstanceResponseRunStatus{
			value: 5,
		}, E_6: DeleteInstanceResponseRunStatus{
			value: 6,
		}, E_7: DeleteInstanceResponseRunStatus{
			value: 7,
		}, E_8: DeleteInstanceResponseRunStatus{
			value: 8,
		},
	}
}

func (c DeleteInstanceResponseRunStatus) Value() int32 {
	return c.value
}

func (c DeleteInstanceResponseRunStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteInstanceResponseRunStatus) UnmarshalJSON(b []byte) error {
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

type DeleteInstanceResponseAccessStatus struct {
	value int32
}

type DeleteInstanceResponseAccessStatusEnum struct {
	E_0 DeleteInstanceResponseAccessStatus
	E_1 DeleteInstanceResponseAccessStatus
}

func GetDeleteInstanceResponseAccessStatusEnum() DeleteInstanceResponseAccessStatusEnum {
	return DeleteInstanceResponseAccessStatusEnum{
		E_0: DeleteInstanceResponseAccessStatus{
			value: 0,
		}, E_1: DeleteInstanceResponseAccessStatus{
			value: 1,
		},
	}
}

func (c DeleteInstanceResponseAccessStatus) Value() int32 {
	return c.value
}

func (c DeleteInstanceResponseAccessStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteInstanceResponseAccessStatus) UnmarshalJSON(b []byte) error {
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
