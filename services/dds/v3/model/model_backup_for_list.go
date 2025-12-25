package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BackupForList 备份信息。
type BackupForList struct {

	// **参数解释：** 备份ID。 **取值范围：** 不涉及。
	Id string `json:"id"`

	// **参数解释：** 备份名称。 **取值范围：** 不涉及。
	Name string `json:"name"`

	// **参数解释：** 备份所属的实例ID，可以调用“查询实例列表”接口获取。如果未申请实例，可以调用“创建实例”接口创建。 **取值范围：** 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释：** 备份所属的实例名称。 **取值范围：** 不涉及。
	InstanceName string `json:"instance_name"`

	Datastore *BackupDatabase `json:"datastore"`

	// **参数解释：** 备份类型。 **取值范围：** - 取值为“Auto”，表示自动全量备份。 - 取值为“Manual”，表示手动全量备份。 - 取值为“Incremental”，表示自动增量备份。
	Type BackupForListType `json:"type"`

	// **参数解释：** 备份开始时间，格式为“yyyy-mm-dd hh:mm:ss”。该时间为UTC时间。 **取值范围：** 不涉及。
	BeginTime string `json:"begin_time"`

	// **参数解释：** 备份结束时间，格式为“yyyy-mm-dd hh:mm:ss”。该时间为UTC时间。 **取值范围：** 不涉及。
	EndTime string `json:"end_time"`

	// **参数解释：** 备份状态。 **取值范围：** - BUILDING：备份中。 - COMPLETED：备份完成。 - FAILED：备份失败。 - DISABLED：备份删除中。
	Status BackupForListStatus `json:"status"`

	// **参数解释：** 备份大小，单位：KB。 **取值范围：** 不涉及。
	Size int64 `json:"size"`

	// **参数解释：** 备份描述。 **取值范围：** 不涉及。
	Description string `json:"description"`

	// **参数解释：** 实例状态。 **取值范围：** - normal，表示实例正常。 - abnormal，表示实例异常。 - creating，表示实例创建中。 - frozen，表示实例被冻结。 - data_disk_full，表示存储空间满。 - enlargefail，表示实例扩容节点个数失败。
	InstanceStatus *BackupForListInstanceStatus `json:"instance_status,omitempty"`

	// **参数解释：** 实例模式。 **取值范围：** - Sharding - ReplicaSet - Single
	InstanceMode *string `json:"instance_mode,omitempty"`

	// **参数解释：** 当前实例是否处于恢复中或者恢复检查中。 **取值范围：** - true，表示实例处于恢复中或者恢复检查中，不允许删除该实例备份。 - false，当前实例未处于恢复中或者恢复检查中，允许删除该实例备份。
	IsInstanceRestoring *bool `json:"is_instance_restoring,omitempty"`

	// **参数解释：** 备份方式。 **取值范围：** - Snapshot，快照备份。 - Physical，物理备份。 - Logical，逻辑备份。 - Incremental，增量备份。
	BackupMethod *string `json:"backup_method,omitempty"`

	// **参数解释：** 是否开启kms加密。 **取值范围：** - true，已开启kms加密。 - false，未开启kms加密。
	KmsEnable *bool `json:"kms_enable,omitempty"`

	// **参数解释：** 是否支持删除该备份。当全备策略存在时，不允许删除自动备份。手动备份允许删除。 **取值范围：** - true，允许删除该备份。 - false，不允许删除该备份。
	Deletable *bool `json:"deletable,omitempty"`
}

func (o BackupForList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupForList struct{}"
	}

	return strings.Join([]string{"BackupForList", string(data)}, " ")
}

type BackupForListType struct {
	value string
}

type BackupForListTypeEnum struct {
	AUTO        BackupForListType
	MANUAL      BackupForListType
	FRAGMENT    BackupForListType
	INCREMENTAL BackupForListType
}

func GetBackupForListTypeEnum() BackupForListTypeEnum {
	return BackupForListTypeEnum{
		AUTO: BackupForListType{
			value: "auto",
		},
		MANUAL: BackupForListType{
			value: "manual",
		},
		FRAGMENT: BackupForListType{
			value: "fragment",
		},
		INCREMENTAL: BackupForListType{
			value: "incremental",
		},
	}
}

func (c BackupForListType) Value() string {
	return c.value
}

func (c BackupForListType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupForListType) UnmarshalJSON(b []byte) error {
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

type BackupForListStatus struct {
	value string
}

type BackupForListStatusEnum struct {
	BUILDING  BackupForListStatus
	COMPLETED BackupForListStatus
	FAILED    BackupForListStatus
	DELETING  BackupForListStatus
}

func GetBackupForListStatusEnum() BackupForListStatusEnum {
	return BackupForListStatusEnum{
		BUILDING: BackupForListStatus{
			value: "BUILDING",
		},
		COMPLETED: BackupForListStatus{
			value: "COMPLETED",
		},
		FAILED: BackupForListStatus{
			value: "FAILED",
		},
		DELETING: BackupForListStatus{
			value: "DELETING",
		},
	}
}

func (c BackupForListStatus) Value() string {
	return c.value
}

func (c BackupForListStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupForListStatus) UnmarshalJSON(b []byte) error {
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

type BackupForListInstanceStatus struct {
	value string
}

type BackupForListInstanceStatusEnum struct {
	NORMAL         BackupForListInstanceStatus
	ABNORMAL       BackupForListInstanceStatus
	CREATING       BackupForListInstanceStatus
	FROZEN         BackupForListInstanceStatus
	DATA_DISK_FULL BackupForListInstanceStatus
	ENLARGEFAIL    BackupForListInstanceStatus
}

func GetBackupForListInstanceStatusEnum() BackupForListInstanceStatusEnum {
	return BackupForListInstanceStatusEnum{
		NORMAL: BackupForListInstanceStatus{
			value: "normal",
		},
		ABNORMAL: BackupForListInstanceStatus{
			value: "abnormal",
		},
		CREATING: BackupForListInstanceStatus{
			value: "creating",
		},
		FROZEN: BackupForListInstanceStatus{
			value: "frozen",
		},
		DATA_DISK_FULL: BackupForListInstanceStatus{
			value: "data_disk_full",
		},
		ENLARGEFAIL: BackupForListInstanceStatus{
			value: "enlargefail",
		},
	}
}

func (c BackupForListInstanceStatus) Value() string {
	return c.value
}

func (c BackupForListInstanceStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupForListInstanceStatus) UnmarshalJSON(b []byte) error {
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
