package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type QueryInstanceBackupResponseBodyBackups struct {

	// 备份ID。
	Id string `json:"id"`

	// 备份名称。
	Name string `json:"name"`

	// 备份描述信息。
	Description string `json:"description"`

	// 备份开始时间，格式为“yyyy-mm-dd hh:mm:ss”。该时间为UTC时间。
	BeginTime *sdktime.SdkTime `json:"begin_time"`

	// 备份结束时间，格式为“yyyy-mm-dd hh:mm:ss”。该时间为UTC时间。
	EndTime *sdktime.SdkTime `json:"end_time"`

	// 备份状态。
	Status QueryInstanceBackupResponseBodyBackupsStatus `json:"status"`

	// 备份大小，单位：KB。
	Size float64 `json:"size"`

	// 备份类型。
	Type QueryInstanceBackupResponseBodyBackupsType `json:"type"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 实例名称。
	InstanceName string `json:"instance_name"`

	Datastore *QueryInstanceBackupResponseBodyDatastore `json:"datastore"`
}

func (o QueryInstanceBackupResponseBodyBackups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryInstanceBackupResponseBodyBackups struct{}"
	}

	return strings.Join([]string{"QueryInstanceBackupResponseBodyBackups", string(data)}, " ")
}

type QueryInstanceBackupResponseBodyBackupsStatus struct {
	value string
}

type QueryInstanceBackupResponseBodyBackupsStatusEnum struct {
	BUILDING  QueryInstanceBackupResponseBodyBackupsStatus
	COMPLETED QueryInstanceBackupResponseBodyBackupsStatus
	FAILED    QueryInstanceBackupResponseBodyBackupsStatus
}

func GetQueryInstanceBackupResponseBodyBackupsStatusEnum() QueryInstanceBackupResponseBodyBackupsStatusEnum {
	return QueryInstanceBackupResponseBodyBackupsStatusEnum{
		BUILDING: QueryInstanceBackupResponseBodyBackupsStatus{
			value: "BUILDING：备份中",
		},
		COMPLETED: QueryInstanceBackupResponseBodyBackupsStatus{
			value: "COMPLETED：备份完成",
		},
		FAILED: QueryInstanceBackupResponseBodyBackupsStatus{
			value: "FAILED：备份失败",
		},
	}
}

func (c QueryInstanceBackupResponseBodyBackupsStatus) Value() string {
	return c.value
}

func (c QueryInstanceBackupResponseBodyBackupsStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryInstanceBackupResponseBodyBackupsStatus) UnmarshalJSON(b []byte) error {
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

type QueryInstanceBackupResponseBodyBackupsType struct {
	value string
}

type QueryInstanceBackupResponseBodyBackupsTypeEnum struct {
	AUTO   QueryInstanceBackupResponseBodyBackupsType
	MANUAL QueryInstanceBackupResponseBodyBackupsType
}

func GetQueryInstanceBackupResponseBodyBackupsTypeEnum() QueryInstanceBackupResponseBodyBackupsTypeEnum {
	return QueryInstanceBackupResponseBodyBackupsTypeEnum{
		AUTO: QueryInstanceBackupResponseBodyBackupsType{
			value: "Auto 自动全量备份",
		},
		MANUAL: QueryInstanceBackupResponseBodyBackupsType{
			value: "Manual 手动全量备份。",
		},
	}
}

func (c QueryInstanceBackupResponseBodyBackupsType) Value() string {
	return c.value
}

func (c QueryInstanceBackupResponseBodyBackupsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryInstanceBackupResponseBodyBackupsType) UnmarshalJSON(b []byte) error {
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
