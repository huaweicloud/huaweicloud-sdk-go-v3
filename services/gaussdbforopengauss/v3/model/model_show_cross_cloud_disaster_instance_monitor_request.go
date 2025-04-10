package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowCrossCloudDisasterInstanceMonitorRequest Request Object
type ShowCrossCloudDisasterInstanceMonitorRequest struct {

	// 语言。
	XLanguage *ShowCrossCloudDisasterInstanceMonitorRequestXLanguage `json:"X-Language,omitempty"`

	// 实例id。
	InstanceId string `json:"instance_id"`

	// 容灾类型
	DisasterType *ShowCrossCloudDisasterInstanceMonitorRequestDisasterType `json:"disaster_type,omitempty"`
}

func (o ShowCrossCloudDisasterInstanceMonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCrossCloudDisasterInstanceMonitorRequest struct{}"
	}

	return strings.Join([]string{"ShowCrossCloudDisasterInstanceMonitorRequest", string(data)}, " ")
}

type ShowCrossCloudDisasterInstanceMonitorRequestXLanguage struct {
	value string
}

type ShowCrossCloudDisasterInstanceMonitorRequestXLanguageEnum struct {
	ZH_CN ShowCrossCloudDisasterInstanceMonitorRequestXLanguage
	EN_US ShowCrossCloudDisasterInstanceMonitorRequestXLanguage
}

func GetShowCrossCloudDisasterInstanceMonitorRequestXLanguageEnum() ShowCrossCloudDisasterInstanceMonitorRequestXLanguageEnum {
	return ShowCrossCloudDisasterInstanceMonitorRequestXLanguageEnum{
		ZH_CN: ShowCrossCloudDisasterInstanceMonitorRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowCrossCloudDisasterInstanceMonitorRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowCrossCloudDisasterInstanceMonitorRequestXLanguage) Value() string {
	return c.value
}

func (c ShowCrossCloudDisasterInstanceMonitorRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowCrossCloudDisasterInstanceMonitorRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ShowCrossCloudDisasterInstanceMonitorRequestDisasterType struct {
	value string
}

type ShowCrossCloudDisasterInstanceMonitorRequestDisasterTypeEnum struct {
	STREAM ShowCrossCloudDisasterInstanceMonitorRequestDisasterType
	DORADO ShowCrossCloudDisasterInstanceMonitorRequestDisasterType
}

func GetShowCrossCloudDisasterInstanceMonitorRequestDisasterTypeEnum() ShowCrossCloudDisasterInstanceMonitorRequestDisasterTypeEnum {
	return ShowCrossCloudDisasterInstanceMonitorRequestDisasterTypeEnum{
		STREAM: ShowCrossCloudDisasterInstanceMonitorRequestDisasterType{
			value: "stream",
		},
		DORADO: ShowCrossCloudDisasterInstanceMonitorRequestDisasterType{
			value: "dorado",
		},
	}
}

func (c ShowCrossCloudDisasterInstanceMonitorRequestDisasterType) Value() string {
	return c.value
}

func (c ShowCrossCloudDisasterInstanceMonitorRequestDisasterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowCrossCloudDisasterInstanceMonitorRequestDisasterType) UnmarshalJSON(b []byte) error {
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
