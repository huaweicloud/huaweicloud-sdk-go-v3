package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateCrossCloudConstructDisasterRequest Request Object
type CreateCrossCloudConstructDisasterRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *CreateCrossCloudConstructDisasterRequestXLanguage `json:"X-Language,omitempty"`

	// 实例id。
	InstanceId string `json:"instance_id"`

	Body *ConstructReq `json:"body,omitempty"`
}

func (o CreateCrossCloudConstructDisasterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCrossCloudConstructDisasterRequest struct{}"
	}

	return strings.Join([]string{"CreateCrossCloudConstructDisasterRequest", string(data)}, " ")
}

type CreateCrossCloudConstructDisasterRequestXLanguage struct {
	value string
}

type CreateCrossCloudConstructDisasterRequestXLanguageEnum struct {
	ZH_CN CreateCrossCloudConstructDisasterRequestXLanguage
	EN_US CreateCrossCloudConstructDisasterRequestXLanguage
}

func GetCreateCrossCloudConstructDisasterRequestXLanguageEnum() CreateCrossCloudConstructDisasterRequestXLanguageEnum {
	return CreateCrossCloudConstructDisasterRequestXLanguageEnum{
		ZH_CN: CreateCrossCloudConstructDisasterRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CreateCrossCloudConstructDisasterRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CreateCrossCloudConstructDisasterRequestXLanguage) Value() string {
	return c.value
}

func (c CreateCrossCloudConstructDisasterRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCrossCloudConstructDisasterRequestXLanguage) UnmarshalJSON(b []byte) error {
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
