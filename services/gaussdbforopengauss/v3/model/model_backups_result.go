package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BackupsResult struct {

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
	Status *BackupsResultStatus `json:"status,omitempty"`

	// 备份大小(单位：MB)
	Size *float64 `json:"size,omitempty"`

	// 备份类型
	Type *BackupsResultType `json:"type,omitempty"`

	Datastore *OpenGaussDatastoreResult `json:"datastore,omitempty"`

	// **参数解释**: 实例ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o BackupsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupsResult struct{}"
	}

	return strings.Join([]string{"BackupsResult", string(data)}, " ")
}

type BackupsResultStatus struct {
	value string
}

type BackupsResultStatusEnum struct {
	BUILDING  BackupsResultStatus
	COMPLETED BackupsResultStatus
	FAILED    BackupsResultStatus
}

func GetBackupsResultStatusEnum() BackupsResultStatusEnum {
	return BackupsResultStatusEnum{
		BUILDING: BackupsResultStatus{
			value: "BUILDING：备份中。",
		},
		COMPLETED: BackupsResultStatus{
			value: "COMPLETED：备份完成。",
		},
		FAILED: BackupsResultStatus{
			value: "FAILED：备份失败。。",
		},
	}
}

func (c BackupsResultStatus) Value() string {
	return c.value
}

func (c BackupsResultStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupsResultStatus) UnmarshalJSON(b []byte) error {
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

type BackupsResultType struct {
	value string
}

type BackupsResultTypeEnum struct {
	AUTO   BackupsResultType
	MANUAL BackupsResultType
}

func GetBackupsResultTypeEnum() BackupsResultTypeEnum {
	return BackupsResultTypeEnum{
		AUTO: BackupsResultType{
			value: "auto：自动全量备份。",
		},
		MANUAL: BackupsResultType{
			value: "manual：手动全量备份。",
		},
	}
}

func (c BackupsResultType) Value() string {
	return c.value
}

func (c BackupsResultType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupsResultType) UnmarshalJSON(b []byte) error {
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
