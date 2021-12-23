package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 查询保护组数据结构
type ShowProtectionGroupParams struct {
	// 保护组的ID。

	Id string `json:"id"`
	// 保护组的名称。

	Name string `json:"name"`
	// 保护组的描述。

	Description string `json:"description"`
	// 保护组的状态。

	Status string `json:"status"`
	// 保护组的同步进度。单位：百分比（%）。

	Progress int32 `json:"progress"`
	// 保护组创建时的生产站点可用区名称。注意：保护组切换、故障切换后，该值不变。

	SourceAvailabilityZone string `json:"source_availability_zone"`
	// 保护组创建时的容灾站点可用区名称。注意：保护组切换、故障切换后，该值不变。

	TargetAvailabilityZone string `json:"target_availability_zone"`
	// 双活域ID。

	DomainId string `json:"domain_id"`
	// 双活域名称。

	DomainName string `json:"domain_name"`
	// 用于标识保护组的当前生产站点。 source：表示当前生产站点可用区为source_availability_zone的值。 target：表示当前生产站点可用区为target_availability_zone的值。

	PriorityStation string `json:"priority_station"`
	// 该保护组中保护实例的个数。

	ProtectedInstanceNum int32 `json:"protected_instance_num"`
	// 该保护组中复制对的个数。

	ReplicationNum int32 `json:"replication_num"`
	// 该保护组中容灾演练的个数。

	DisasterRecoveryDrillNum int32 `json:"disaster_recovery_drill_num"`
	// 保护状态。started：表示该保护组开始保护。stopped：表示该保护组停止保护。 说明:系统近期进行了升级，对于升级后创建的保护组，该字段值为null，无实际意义。

	ProtectedStatus ShowProtectionGroupParamsProtectedStatus `json:"protected_status"`
	// 数据同步状态。 active：表示数据已同步完成。 inactive：表示数据未同步。 copying：表示数据正在同步。 active-stopped：表示数据已停止同步。  说明:系统近期进行了升级，对于升级后创建的保护组，该字段值为null，无实际意义。

	ReplicationStatus ShowProtectionGroupParamsReplicationStatus `json:"replication_status"`
	// 健康状态。 normal：表示该保护组处于正常状态。 abnormal：表示该保护组处于非正常状态。  说明:系统近期进行了升级，对于升级后创建的保护组，该字段值为null，无实际意义。

	HealthStatus ShowProtectionGroupParamsHealthStatus `json:"health_status"`
	// 生产站点虚拟私有云ID。

	SourceVpcId string `json:"source_vpc_id"`
	// 容灾站点虚拟私有云ID。

	TargetVpcId string `json:"target_vpc_id"`
	// 容灾演练虚拟私有云ID。（该参数暂未使用）

	TestVpcId string `json:"test_vpc_id"`
	// 部署模式。默认值为“migration”，migration表示VPC内迁移。

	DrType string `json:"dr_type"`
	// 创建时间。默认格式为：\"yyyy-MM-dd'T'HH:mm:ss.SSSZ\"，例：\"2019-04-01T12:00:00.000Z\"。

	CreatedAt string `json:"created_at"`
	// 更新时间。默认格式为：\"yyyy-MM-dd'T'HH:mm:ss.SSSZ\"，例：\"2019-04-01T12:00:00.000Z\"。

	UpdatedAt string `json:"updated_at"`
	// 保护模式。 replication-pair：表示以复制对为单位进行数据同步。 null：表示将保护组中的所有复制对作为一个整体进行数据同步。  说明:当保护组中的所有复制对作为一个整体进行数据同步时，如果数据同步失败，保护组中的所有复制对都会受到影响。因此，SDRS服务对系统做了优化升级： 对于升级后创建的资源，默认以复制对为单位进行数据同步，返回值为replication-pair； 对于已有资源，仍以一个整体进行数据同步，返回值为null。

	ProtectionType ShowProtectionGroupParamsProtectionType `json:"protection_type"`
	// 复制类型。 说明:预留参数，暂未启用。

	ReplicationModel string `json:"replication_model"`
	// 管理的服务器类型 ECS：表示管理的服务器类型为云服务器。

	ServerType string `json:"server_type"`
}

func (o ShowProtectionGroupParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProtectionGroupParams struct{}"
	}

	return strings.Join([]string{"ShowProtectionGroupParams", string(data)}, " ")
}

type ShowProtectionGroupParamsProtectedStatus struct {
	value string
}

type ShowProtectionGroupParamsProtectedStatusEnum struct {
	STARTED ShowProtectionGroupParamsProtectedStatus
	STOPPED ShowProtectionGroupParamsProtectedStatus
}

func GetShowProtectionGroupParamsProtectedStatusEnum() ShowProtectionGroupParamsProtectedStatusEnum {
	return ShowProtectionGroupParamsProtectedStatusEnum{
		STARTED: ShowProtectionGroupParamsProtectedStatus{
			value: "started",
		},
		STOPPED: ShowProtectionGroupParamsProtectedStatus{
			value: "stopped",
		},
	}
}

func (c ShowProtectionGroupParamsProtectedStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowProtectionGroupParamsProtectedStatus) UnmarshalJSON(b []byte) error {
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

type ShowProtectionGroupParamsReplicationStatus struct {
	value string
}

type ShowProtectionGroupParamsReplicationStatusEnum struct {
	ACTIVE         ShowProtectionGroupParamsReplicationStatus
	INACTIVE       ShowProtectionGroupParamsReplicationStatus
	COPYING        ShowProtectionGroupParamsReplicationStatus
	ACTIVE_STOPPED ShowProtectionGroupParamsReplicationStatus
}

func GetShowProtectionGroupParamsReplicationStatusEnum() ShowProtectionGroupParamsReplicationStatusEnum {
	return ShowProtectionGroupParamsReplicationStatusEnum{
		ACTIVE: ShowProtectionGroupParamsReplicationStatus{
			value: "active",
		},
		INACTIVE: ShowProtectionGroupParamsReplicationStatus{
			value: "inactive",
		},
		COPYING: ShowProtectionGroupParamsReplicationStatus{
			value: "copying",
		},
		ACTIVE_STOPPED: ShowProtectionGroupParamsReplicationStatus{
			value: "active-stopped",
		},
	}
}

func (c ShowProtectionGroupParamsReplicationStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowProtectionGroupParamsReplicationStatus) UnmarshalJSON(b []byte) error {
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

type ShowProtectionGroupParamsHealthStatus struct {
	value string
}

type ShowProtectionGroupParamsHealthStatusEnum struct {
	NORMAL   ShowProtectionGroupParamsHealthStatus
	ABNORMAL ShowProtectionGroupParamsHealthStatus
	NULL     ShowProtectionGroupParamsHealthStatus
}

func GetShowProtectionGroupParamsHealthStatusEnum() ShowProtectionGroupParamsHealthStatusEnum {
	return ShowProtectionGroupParamsHealthStatusEnum{
		NORMAL: ShowProtectionGroupParamsHealthStatus{
			value: "normal",
		},
		ABNORMAL: ShowProtectionGroupParamsHealthStatus{
			value: "abnormal",
		},
		NULL: ShowProtectionGroupParamsHealthStatus{
			value: "null",
		},
	}
}

func (c ShowProtectionGroupParamsHealthStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowProtectionGroupParamsHealthStatus) UnmarshalJSON(b []byte) error {
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

type ShowProtectionGroupParamsProtectionType struct {
	value string
}

type ShowProtectionGroupParamsProtectionTypeEnum struct {
	REPLICATION_PAIR ShowProtectionGroupParamsProtectionType
	NULL             ShowProtectionGroupParamsProtectionType
}

func GetShowProtectionGroupParamsProtectionTypeEnum() ShowProtectionGroupParamsProtectionTypeEnum {
	return ShowProtectionGroupParamsProtectionTypeEnum{
		REPLICATION_PAIR: ShowProtectionGroupParamsProtectionType{
			value: "replication-pair",
		},
		NULL: ShowProtectionGroupParamsProtectionType{
			value: "null",
		},
	}
}

func (c ShowProtectionGroupParamsProtectionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowProtectionGroupParamsProtectionType) UnmarshalJSON(b []byte) error {
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
