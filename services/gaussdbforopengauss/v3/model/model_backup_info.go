package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 备份信息。
type BackupInfo struct {

	// 备份ID。
	Id string `json:"id" xml:"id"`

	// 备份名称。
	Name string `json:"name" xml:"name"`

	// 备份描述。
	Description string `json:"description" xml:"description"`

	// 备份开始时间，格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	BeginTime string `json:"begin_time" xml:"begin_time"`

	// 备份状态，取值：  - BUILDING: 备份中。 - COMPLETED: 备份完成。 - FAILED：备份失败。 - DELETING：备份删除中。
	Status BackupInfoStatus `json:"status" xml:"status"`

	// 备份类型，取值： - “manual”: 手动全量备份
	Type BackupInfoType `json:"type" xml:"type"`

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`
}

func (o BackupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupInfo struct{}"
	}

	return strings.Join([]string{"BackupInfo", string(data)}, " ")
}

type BackupInfoStatus struct {
	value string
}

type BackupInfoStatusEnum struct {
	BUILDING  BackupInfoStatus
	COMPLETED BackupInfoStatus
	FAILED    BackupInfoStatus
	DELETING  BackupInfoStatus
}

func GetBackupInfoStatusEnum() BackupInfoStatusEnum {
	return BackupInfoStatusEnum{
		BUILDING: BackupInfoStatus{
			value: "BUILDING",
		},
		COMPLETED: BackupInfoStatus{
			value: "COMPLETED",
		},
		FAILED: BackupInfoStatus{
			value: "FAILED",
		},
		DELETING: BackupInfoStatus{
			value: "DELETING",
		},
	}
}

func (c BackupInfoStatus) Value() string {
	return c.value
}

func (c BackupInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupInfoStatus) UnmarshalJSON(b []byte) error {
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

type BackupInfoType struct {
	value string
}

type BackupInfoTypeEnum struct {
	MANUAL BackupInfoType
}

func GetBackupInfoTypeEnum() BackupInfoTypeEnum {
	return BackupInfoTypeEnum{
		MANUAL: BackupInfoType{
			value: "manual",
		},
	}
}

func (c BackupInfoType) Value() string {
	return c.value
}

func (c BackupInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupInfoType) UnmarshalJSON(b []byte) error {
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
