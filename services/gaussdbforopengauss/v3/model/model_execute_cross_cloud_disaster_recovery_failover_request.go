package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExecuteCrossCloudDisasterRecoveryFailoverRequest Request Object
type ExecuteCrossCloudDisasterRecoveryFailoverRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *ExecuteCrossCloudDisasterRecoveryFailoverRequestXLanguage `json:"X-Language,omitempty"`

	// 实例id。
	InstanceId string `json:"instance_id"`

	Body *DisasterFailoverReqBody `json:"body,omitempty"`
}

func (o ExecuteCrossCloudDisasterRecoveryFailoverRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteCrossCloudDisasterRecoveryFailoverRequest struct{}"
	}

	return strings.Join([]string{"ExecuteCrossCloudDisasterRecoveryFailoverRequest", string(data)}, " ")
}

type ExecuteCrossCloudDisasterRecoveryFailoverRequestXLanguage struct {
	value string
}

type ExecuteCrossCloudDisasterRecoveryFailoverRequestXLanguageEnum struct {
	ZH_CN ExecuteCrossCloudDisasterRecoveryFailoverRequestXLanguage
	EN_US ExecuteCrossCloudDisasterRecoveryFailoverRequestXLanguage
}

func GetExecuteCrossCloudDisasterRecoveryFailoverRequestXLanguageEnum() ExecuteCrossCloudDisasterRecoveryFailoverRequestXLanguageEnum {
	return ExecuteCrossCloudDisasterRecoveryFailoverRequestXLanguageEnum{
		ZH_CN: ExecuteCrossCloudDisasterRecoveryFailoverRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ExecuteCrossCloudDisasterRecoveryFailoverRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ExecuteCrossCloudDisasterRecoveryFailoverRequestXLanguage) Value() string {
	return c.value
}

func (c ExecuteCrossCloudDisasterRecoveryFailoverRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteCrossCloudDisasterRecoveryFailoverRequestXLanguage) UnmarshalJSON(b []byte) error {
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
