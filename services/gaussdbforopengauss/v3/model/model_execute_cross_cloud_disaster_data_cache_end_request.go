package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExecuteCrossCloudDisasterDataCacheEndRequest Request Object
type ExecuteCrossCloudDisasterDataCacheEndRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *ExecuteCrossCloudDisasterDataCacheEndRequestXLanguage `json:"X-Language,omitempty"`

	// 实例id。
	InstanceId string `json:"instance_id"`

	Body *DisasterRecoverStopXlogKeepRequestBody `json:"body,omitempty"`
}

func (o ExecuteCrossCloudDisasterDataCacheEndRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteCrossCloudDisasterDataCacheEndRequest struct{}"
	}

	return strings.Join([]string{"ExecuteCrossCloudDisasterDataCacheEndRequest", string(data)}, " ")
}

type ExecuteCrossCloudDisasterDataCacheEndRequestXLanguage struct {
	value string
}

type ExecuteCrossCloudDisasterDataCacheEndRequestXLanguageEnum struct {
	ZH_CN ExecuteCrossCloudDisasterDataCacheEndRequestXLanguage
	EN_US ExecuteCrossCloudDisasterDataCacheEndRequestXLanguage
}

func GetExecuteCrossCloudDisasterDataCacheEndRequestXLanguageEnum() ExecuteCrossCloudDisasterDataCacheEndRequestXLanguageEnum {
	return ExecuteCrossCloudDisasterDataCacheEndRequestXLanguageEnum{
		ZH_CN: ExecuteCrossCloudDisasterDataCacheEndRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ExecuteCrossCloudDisasterDataCacheEndRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ExecuteCrossCloudDisasterDataCacheEndRequestXLanguage) Value() string {
	return c.value
}

func (c ExecuteCrossCloudDisasterDataCacheEndRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteCrossCloudDisasterDataCacheEndRequestXLanguage) UnmarshalJSON(b []byte) error {
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
