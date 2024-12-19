package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSourceInstanceDetailRequest Request Object
type ShowSourceInstanceDetailRequest struct {

	// 语言。默认值：en-us。
	XLanguage *ShowSourceInstanceDetailRequestXLanguage `json:"X-Language,omitempty"`

	// 原实例ID。  (instance_id 、restore_time为一组)
	InstanceId *string `json:"instance_id,omitempty"`

	// UNIX时间戳格式，单位是毫秒，时区是UTC，某时间点实例的信息。  (instance_id 、restore_time为一组)
	RestoreTime *string `json:"restore_time,omitempty"`

	// 备份ID。  (backup_id为一组)  备份ID不为空时，可以不需要实例ID和时间戳。
	BackupId *string `json:"backup_id,omitempty"`
}

func (o ShowSourceInstanceDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSourceInstanceDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowSourceInstanceDetailRequest", string(data)}, " ")
}

type ShowSourceInstanceDetailRequestXLanguage struct {
	value string
}

type ShowSourceInstanceDetailRequestXLanguageEnum struct {
	ZH_CN ShowSourceInstanceDetailRequestXLanguage
	EN_US ShowSourceInstanceDetailRequestXLanguage
}

func GetShowSourceInstanceDetailRequestXLanguageEnum() ShowSourceInstanceDetailRequestXLanguageEnum {
	return ShowSourceInstanceDetailRequestXLanguageEnum{
		ZH_CN: ShowSourceInstanceDetailRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowSourceInstanceDetailRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowSourceInstanceDetailRequestXLanguage) Value() string {
	return c.value
}

func (c ShowSourceInstanceDetailRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSourceInstanceDetailRequestXLanguage) UnmarshalJSON(b []byte) error {
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
