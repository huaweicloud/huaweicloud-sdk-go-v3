package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Backups struct {

	// 备份ID。
	Id *string `json:"id,omitempty"`

	// 备份名称。
	Name *string `json:"name,omitempty"`

	// 备份文件描述信息。
	Description *string `json:"description,omitempty"`

	// 备份开始时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	BeginTime *string `json:"begin_time,omitempty"`

	// 备份结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	EndTime *string `json:"end_time,omitempty"`

	// 备份状态
	Status *BackupsStatus `json:"status,omitempty"`

	// 备份大小(单位：MB)
	Size *float64 `json:"size,omitempty"`

	// 备份类型
	Type *BackupsType `json:"type,omitempty"`

	Datastore *OpenGaussDatastore `json:"datastore,omitempty"`

	// **参数解释**: 实例ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o Backups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Backups struct{}"
	}

	return strings.Join([]string{"Backups", string(data)}, " ")
}

type BackupsStatus struct {
	value string
}

type BackupsStatusEnum struct {
	BUILDING  BackupsStatus
	COMPLETED BackupsStatus
	FAILED    BackupsStatus
}

func GetBackupsStatusEnum() BackupsStatusEnum {
	return BackupsStatusEnum{
		BUILDING: BackupsStatus{
			value: "BUILDING：备份中。",
		},
		COMPLETED: BackupsStatus{
			value: "COMPLETED：备份完成。",
		},
		FAILED: BackupsStatus{
			value: "FAILED：备份失败。。",
		},
	}
}

func (c BackupsStatus) Value() string {
	return c.value
}

func (c BackupsStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupsStatus) UnmarshalJSON(b []byte) error {
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

type BackupsType struct {
	value string
}

type BackupsTypeEnum struct {
	AUTO   BackupsType
	MANUAL BackupsType
}

func GetBackupsTypeEnum() BackupsTypeEnum {
	return BackupsTypeEnum{
		AUTO: BackupsType{
			value: "auto：自动全量备份。",
		},
		MANUAL: BackupsType{
			value: "manual：手动全量备份。",
		},
	}
}

func (c BackupsType) Value() string {
	return c.value
}

func (c BackupsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupsType) UnmarshalJSON(b []byte) error {
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
