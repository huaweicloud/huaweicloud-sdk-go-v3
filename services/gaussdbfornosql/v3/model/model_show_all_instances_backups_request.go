package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAllInstancesBackupsRequest Request Object
type ShowAllInstancesBackupsRequest struct {

	// 分页页码。
	Offset int32 `json:"offset"`

	// 每页条数。
	Limit int32 `json:"limit"`

	// 引擎类型 不传该参数则查询所有的引擎。
	DatastoreType *ShowAllInstancesBackupsRequestDatastoreType `json:"datastore_type,omitempty"`

	// 实例ID 不传该参数则查询所有备份列表。
	InstanceId *string `json:"instance_id,omitempty"`

	// 备份ID。
	BackupId *string `json:"backup_id,omitempty"`

	// 备份类型。
	BackupType *ShowAllInstancesBackupsRequestBackupType `json:"backup_type,omitempty"`

	// 查询备份开始的时间，格式为“yyyy-mm-dd hh:mm:ss”。该时间为UTC时间。
	BeginTime *string `json:"begin_time,omitempty"`

	// 查询备份开始的结束时间，格式为“yyyy-mm-dd hh:mm:ss”。该时间为UTC时间。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ShowAllInstancesBackupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAllInstancesBackupsRequest struct{}"
	}

	return strings.Join([]string{"ShowAllInstancesBackupsRequest", string(data)}, " ")
}

type ShowAllInstancesBackupsRequestDatastoreType struct {
	value string
}

type ShowAllInstancesBackupsRequestDatastoreTypeEnum struct {
	CASSANDRA ShowAllInstancesBackupsRequestDatastoreType
	MONGODB   ShowAllInstancesBackupsRequestDatastoreType
	REDIS     ShowAllInstancesBackupsRequestDatastoreType
	INFLUXDB  ShowAllInstancesBackupsRequestDatastoreType
}

func GetShowAllInstancesBackupsRequestDatastoreTypeEnum() ShowAllInstancesBackupsRequestDatastoreTypeEnum {
	return ShowAllInstancesBackupsRequestDatastoreTypeEnum{
		CASSANDRA: ShowAllInstancesBackupsRequestDatastoreType{
			value: "cassandra",
		},
		MONGODB: ShowAllInstancesBackupsRequestDatastoreType{
			value: "mongodb",
		},
		REDIS: ShowAllInstancesBackupsRequestDatastoreType{
			value: "redis",
		},
		INFLUXDB: ShowAllInstancesBackupsRequestDatastoreType{
			value: "influxdb",
		},
	}
}

func (c ShowAllInstancesBackupsRequestDatastoreType) Value() string {
	return c.value
}

func (c ShowAllInstancesBackupsRequestDatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAllInstancesBackupsRequestDatastoreType) UnmarshalJSON(b []byte) error {
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

type ShowAllInstancesBackupsRequestBackupType struct {
	value string
}

type ShowAllInstancesBackupsRequestBackupTypeEnum struct {
	AUTO   ShowAllInstancesBackupsRequestBackupType
	MANUAL ShowAllInstancesBackupsRequestBackupType
}

func GetShowAllInstancesBackupsRequestBackupTypeEnum() ShowAllInstancesBackupsRequestBackupTypeEnum {
	return ShowAllInstancesBackupsRequestBackupTypeEnum{
		AUTO: ShowAllInstancesBackupsRequestBackupType{
			value: "Auto 自动全量备份",
		},
		MANUAL: ShowAllInstancesBackupsRequestBackupType{
			value: "Manual 手动全量备份。",
		},
	}
}

func (c ShowAllInstancesBackupsRequestBackupType) Value() string {
	return c.value
}

func (c ShowAllInstancesBackupsRequestBackupType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAllInstancesBackupsRequestBackupType) UnmarshalJSON(b []byte) error {
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
