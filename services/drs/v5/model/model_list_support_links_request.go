package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSupportLinksRequest Request Object
type ListSupportLinksRequest struct {

	// 请求语言类型。
	XLanguage *ListSupportLinksRequestXLanguage `json:"X-Language,omitempty"`

	// 任务场景。取值： - migration：实时迁移。 - sync：实时同步。 - cloudDataGuard：实时灾备。 - replay：录制回放。 - verify：校验任务。 - cdc：CDC任务。
	JobType string `json:"job_type"`

	// 偏移量，表示查询该偏移量后面的记录，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询返回记录的数量限制，默认为10。
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
