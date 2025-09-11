package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExecuteCrossCloudDisasterDataCacheStartRequest Request Object
type ExecuteCrossCloudDisasterDataCacheStartRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *ExecuteCrossCloudDisasterDataCacheStartRequestXLanguage `json:"X-Language,omitempty"`

	// 实例id。
	InstanceId string `json:"instance_id"`

	Body *DisasterRecoverStartXlogKeepRequestBody `json:"body,omitempty"`
}

func (o ExecuteCrossCloudDisasterDataCacheStartRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteCrossCloudDisasterDataCacheStartRequest struct{}"
	}

	return strings.Join([]string{"ExecuteCrossCloudDisasterDataCacheStartRequest", string(data)}, " ")
}

type ExecuteCrossCloudDisasterDataCacheStartRequestXLanguage struct {
	value string
}

type ExecuteCrossCloudDisasterDataCacheStartRequestXLanguageEnum struct {
	ZH_CN ExecuteCrossCloudDisasterDataCacheStartRequestXLanguage
	EN_US ExecuteCrossCloudDisasterDataCacheStartRequestXLanguage
}

func GetExecuteCrossCloudDisasterDataCacheStartRequestXLanguageEnum() ExecuteCrossCloudDisasterDataCacheStartRequestXLanguageEnum {
	return ExecuteCrossCloudDisasterDataCacheStartRequestXLanguageEnum{
		ZH_CN: ExecuteCrossCloudDisasterDataCacheStartRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ExecuteCrossCloudDisasterDataCacheStartRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ExecuteCrossCloudDisasterDataCacheStartRequestXLanguage) Value() string {
	return c.value
}

func (c ExecuteCrossCloudDisasterDataCacheStartRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteCrossCloudDisasterDataCacheStartRequestXLanguage) UnmarshalJSON(b []byte) error {
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
