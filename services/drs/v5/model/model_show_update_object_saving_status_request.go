package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowUpdateObjectSavingStatusRequest Request Object
type ShowUpdateObjectSavingStatusRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ShowUpdateObjectSavingStatusRequestXLanguage `json:"X-Language,omitempty"`

	// 保存对象接口返回的ID。
	QueryId string `json:"query_id"`
}

func (o ShowUpdateObjectSavingStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUpdateObjectSavingStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowUpdateObjectSavingStatusRequest", string(data)}, " ")
}

type ShowUpdateObjectSavingStatusRequestXLanguage struct {
	value string
}

type ShowUpdateObjectSavingStatusRequestXLanguageEnum struct {
	EN_US ShowUpdateObjectSavingStatusRequestXLanguage
	ZH_CN ShowUpdateObjectSavingStatusRequestXLanguage
}

func GetShowUpdateObjectSavingStatusRequestXLanguageEnum() ShowUpdateObjectSavingStatusRequestXLanguageEnum {
	return ShowUpdateObjectSavingStatusRequestXLanguageEnum{
		EN_US: ShowUpdateObjectSavingStatusRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowUpdateObjectSavingStatusRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowUpdateObjectSavingStatusRequestXLanguage) Value() string {
	return c.value
}

func (c ShowUpdateObjectSavingStatusRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowUpdateObjectSavingStatusRequestXLanguage) UnmarshalJSON(b []byte) error {
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
