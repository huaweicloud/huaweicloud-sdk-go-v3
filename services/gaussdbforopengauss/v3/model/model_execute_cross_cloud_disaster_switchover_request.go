package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExecuteCrossCloudDisasterSwitchoverRequest Request Object
type ExecuteCrossCloudDisasterSwitchoverRequest struct {

	// 语言。
	XLanguage *ExecuteCrossCloudDisasterSwitchoverRequestXLanguage `json:"X-Language,omitempty"`

	// 实例id。
	InstanceId string `json:"instance_id"`

	Body *SwitchoverReq `json:"body,omitempty"`
}

func (o ExecuteCrossCloudDisasterSwitchoverRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteCrossCloudDisasterSwitchoverRequest struct{}"
	}

	return strings.Join([]string{"ExecuteCrossCloudDisasterSwitchoverRequest", string(data)}, " ")
}

type ExecuteCrossCloudDisasterSwitchoverRequestXLanguage struct {
	value string
}

type ExecuteCrossCloudDisasterSwitchoverRequestXLanguageEnum struct {
	ZH_CN ExecuteCrossCloudDisasterSwitchoverRequestXLanguage
	EN_US ExecuteCrossCloudDisasterSwitchoverRequestXLanguage
}

func GetExecuteCrossCloudDisasterSwitchoverRequestXLanguageEnum() ExecuteCrossCloudDisasterSwitchoverRequestXLanguageEnum {
	return ExecuteCrossCloudDisasterSwitchoverRequestXLanguageEnum{
		ZH_CN: ExecuteCrossCloudDisasterSwitchoverRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ExecuteCrossCloudDisasterSwitchoverRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ExecuteCrossCloudDisasterSwitchoverRequestXLanguage) Value() string {
	return c.value
}

func (c ExecuteCrossCloudDisasterSwitchoverRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteCrossCloudDisasterSwitchoverRequestXLanguage) UnmarshalJSON(b []byte) error {
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
