package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExecuteCrossCloudReleaseDisasterRequest Request Object
type ExecuteCrossCloudReleaseDisasterRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *ExecuteCrossCloudReleaseDisasterRequestXLanguage `json:"X-Language,omitempty"`

	// 实例id。
	InstanceId string `json:"instance_id"`

	Body *ReleaseDisasterReq `json:"body,omitempty"`
}

func (o ExecuteCrossCloudReleaseDisasterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteCrossCloudReleaseDisasterRequest struct{}"
	}

	return strings.Join([]string{"ExecuteCrossCloudReleaseDisasterRequest", string(data)}, " ")
}

type ExecuteCrossCloudReleaseDisasterRequestXLanguage struct {
	value string
}

type ExecuteCrossCloudReleaseDisasterRequestXLanguageEnum struct {
	ZH_CN ExecuteCrossCloudReleaseDisasterRequestXLanguage
	EN_US ExecuteCrossCloudReleaseDisasterRequestXLanguage
}

func GetExecuteCrossCloudReleaseDisasterRequestXLanguageEnum() ExecuteCrossCloudReleaseDisasterRequestXLanguageEnum {
	return ExecuteCrossCloudReleaseDisasterRequestXLanguageEnum{
		ZH_CN: ExecuteCrossCloudReleaseDisasterRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ExecuteCrossCloudReleaseDisasterRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ExecuteCrossCloudReleaseDisasterRequestXLanguage) Value() string {
	return c.value
}

func (c ExecuteCrossCloudReleaseDisasterRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteCrossCloudReleaseDisasterRequestXLanguage) UnmarshalJSON(b []byte) error {
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
