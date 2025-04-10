package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DisasterRelations 对端容灾关系信息。
type DisasterRelations struct {

	// 容灾类型。
	DisasterType *string `json:"disaster_type,omitempty"`

	// 容灾任务名称。
	Name *string `json:"name,omitempty"`

	// 容灾角色。
	DisasterRole *DisasterRelationsDisasterRole `json:"disaster_role,omitempty"`

	// 创建时间，格式为“yyyy-mm-dd hh:mm:ss”。
	Created *string `json:"created,omitempty"`

	// 更新时间，格式为“yyyy-mm-dd hh:mm:ss”。
	Updated *string `json:"updated,omitempty"`

	SlaveRegionInstanceInfo *RegionInstanceInfo `json:"slave_region_instance_info,omitempty"`

	MasterRegionInstanceInfo *RegionInstanceInfo `json:"master_region_instance_info,omitempty"`

	// 容灾记录id。
	Id *string `json:"id,omitempty"`

	// 容灾关系唯一id。
	SynchronizationId *string `json:"synchronization_id,omitempty"`

	// 当前region实例id。
	InstanceId *string `json:"instance_id,omitempty"`

	// 当前region实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 当前region实例状态。
	InstanceStatus *string `json:"instance_status,omitempty"`

	// 容灾记录状态。
	Status *string `json:"status,omitempty"`

	// 预校验失败原因。
	PrecheckFailedReason *string `json:"precheck_failed_reason,omitempty"`

	// 实例当前操作action。
	Actions *[]string `json:"actions,omitempty"`
}

func (o DisasterRelations) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisasterRelations struct{}"
	}

	return strings.Join([]string{"DisasterRelations", string(data)}, " ")
}

type DisasterRelationsDisasterRole struct {
	value string
}

type DisasterRelationsDisasterRoleEnum struct {
	DISASTER DisasterRelationsDisasterRole
	MASTER   DisasterRelationsDisasterRole
}

func GetDisasterRelationsDisasterRoleEnum() DisasterRelationsDisasterRoleEnum {
	return DisasterRelationsDisasterRoleEnum{
		DISASTER: DisasterRelationsDisasterRole{
			value: "disaster",
		},
		MASTER: DisasterRelationsDisasterRole{
			value: "master",
		},
	}
}

func (c DisasterRelationsDisasterRole) Value() string {
	return c.value
}

func (c DisasterRelationsDisasterRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DisasterRelationsDisasterRole) UnmarshalJSON(b []byte) error {
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
