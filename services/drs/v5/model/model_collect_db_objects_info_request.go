package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CollectDbObjectsInfoRequest Request Object
type CollectDbObjectsInfoRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *CollectDbObjectsInfoRequestXLanguage `json:"X-Language,omitempty"`

	Body *QuerySelectObjectInfoReq `json:"body,omitempty"`
}

func (o CollectDbObjectsInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectDbObjectsInfoRequest struct{}"
	}

	return strings.Join([]string{"CollectDbObjectsInfoRequest", string(data)}, " ")
}

type CollectDbObjectsInfoRequestXLanguage struct {
	value string
}

type CollectDbObjectsInfoRequestXLanguageEnum struct {
	EN_US CollectDbObjectsInfoRequestXLanguage
	ZH_CN CollectDbObjectsInfoRequestXLanguage
}

func GetCollectDbObjectsInfoRequestXLanguageEnum() CollectDbObjectsInfoRequestXLanguageEnum {
	return CollectDbObjectsInfoRequestXLanguageEnum{
		EN_US: CollectDbObjectsInfoRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CollectDbObjectsInfoRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c CollectDbObjectsInfoRequestXLanguage) Value() string {
	return c.value
}

func (c CollectDbObjectsInfoRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CollectDbObjectsInfoRequestXLanguage) UnmarshalJSON(b []byte) error {
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
