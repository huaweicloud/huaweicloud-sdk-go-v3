package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSupportLinksRequest Request Object
type ListSupportLinksRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us：英文 - zh-cn：中文  **默认取值**：  en-us
	XLanguage *ListSupportLinksRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**：  DRS任务场景类型，用于区分不同的数据复制场景。  **约束限制**：  不涉及。  **取值范围**：  - migration：实时迁移 - sync：实时同步 - cloudDataGuard：实时灾备 - replay：录制回放 - verify：校验任务 - cdc：CDC任务  **默认取值**：  不涉及
	JobType string `json:"job_type"`

	// **参数解释**：  偏移量，表示查询该偏移量后面的记录，用于分页查询。  **约束限制**：  最大为10000。  **取值范围**：  不涉及。  **默认取值**：  0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：  查询返回记录的数量限制，用于分页查询时控制每页返回的数据条数。  **约束限制**：  最大为2000。  **取值范围**：  不涉及。  **默认取值**：  10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSupportLinksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportLinksRequest struct{}"
	}

	return strings.Join([]string{"ListSupportLinksRequest", string(data)}, " ")
}

type ListSupportLinksRequestXLanguage struct {
	value string
}

type ListSupportLinksRequestXLanguageEnum struct {
	EN_US ListSupportLinksRequestXLanguage
	ZH_CN ListSupportLinksRequestXLanguage
}

func GetListSupportLinksRequestXLanguageEnum() ListSupportLinksRequestXLanguageEnum {
	return ListSupportLinksRequestXLanguageEnum{
		EN_US: ListSupportLinksRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListSupportLinksRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListSupportLinksRequestXLanguage) Value() string {
	return c.value
}

func (c ListSupportLinksRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSupportLinksRequestXLanguage) UnmarshalJSON(b []byte) error {
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
