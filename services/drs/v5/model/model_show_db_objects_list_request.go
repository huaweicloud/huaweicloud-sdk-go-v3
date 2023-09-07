package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDbObjectsListRequest Request Object
type ShowDbObjectsListRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ShowDbObjectsListRequestXLanguage `json:"X-Language,omitempty"`

	// 查询对象信息类型。 取值：  - modified：查询已选择的（已同步的和未下发的）对象信息。
	Type string `json:"type"`
}

func (o ShowDbObjectsListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDbObjectsListRequest struct{}"
	}

	return strings.Join([]string{"ShowDbObjectsListRequest", string(data)}, " ")
}

type ShowDbObjectsListRequestXLanguage struct {
	value string
}

type ShowDbObjectsListRequestXLanguageEnum struct {
	EN_US ShowDbObjectsListRequestXLanguage
	ZH_CN ShowDbObjectsListRequestXLanguage
}

func GetShowDbObjectsListRequestXLanguageEnum() ShowDbObjectsListRequestXLanguageEnum {
	return ShowDbObjectsListRequestXLanguageEnum{
		EN_US: ShowDbObjectsListRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowDbObjectsListRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowDbObjectsListRequestXLanguage) Value() string {
	return c.value
}

func (c ShowDbObjectsListRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDbObjectsListRequestXLanguage) UnmarshalJSON(b []byte) error {
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
