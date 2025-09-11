package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExecuteCrossCloudDisasterRestoreRequest Request Object
type ExecuteCrossCloudDisasterRestoreRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *ExecuteCrossCloudDisasterRestoreRequestXLanguage `json:"X-Language,omitempty"`

	// 实例id。
	InstanceId string `json:"instance_id"`

	Body *RestoreDisasterReq `json:"body,omitempty"`
}

func (o ExecuteCrossCloudDisasterRestoreRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteCrossCloudDisasterRestoreRequest struct{}"
	}

	return strings.Join([]string{"ExecuteCrossCloudDisasterRestoreRequest", string(data)}, " ")
}

type ExecuteCrossCloudDisasterRestoreRequestXLanguage struct {
	value string
}

type ExecuteCrossCloudDisasterRestoreRequestXLanguageEnum struct {
	ZH_CN ExecuteCrossCloudDisasterRestoreRequestXLanguage
	EN_US ExecuteCrossCloudDisasterRestoreRequestXLanguage
}

func GetExecuteCrossCloudDisasterRestoreRequestXLanguageEnum() ExecuteCrossCloudDisasterRestoreRequestXLanguageEnum {
	return ExecuteCrossCloudDisasterRestoreRequestXLanguageEnum{
		ZH_CN: ExecuteCrossCloudDisasterRestoreRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ExecuteCrossCloudDisasterRestoreRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ExecuteCrossCloudDisasterRestoreRequestXLanguage) Value() string {
	return c.value
}

func (c ExecuteCrossCloudDisasterRestoreRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteCrossCloudDisasterRestoreRequestXLanguage) UnmarshalJSON(b []byte) error {
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
