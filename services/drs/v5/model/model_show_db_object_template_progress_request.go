package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDbObjectTemplateProgressRequest Request Object
type ShowDbObjectTemplateProgressRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ShowDbObjectTemplateProgressRequestXLanguage `json:"X-Language,omitempty"`

	// 偏移量，表示查询该偏移量后面的记录。
	Offset *int32 `json:"offset,omitempty"`

	// 查询返回记录的数量限制。
	Limit *int32 `json:"limit,omitempty"`

	// 默认为空。 - column：当进行列加工导入时，查询列加工导入进度。
	Type *string `json:"type,omitempty"`
}

func (o ShowDbObjectTemplateProgressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDbObjectTemplateProgressRequest struct{}"
	}

	return strings.Join([]string{"ShowDbObjectTemplateProgressRequest", string(data)}, " ")
}

type ShowDbObjectTemplateProgressRequestXLanguage struct {
	value string
}

type ShowDbObjectTemplateProgressRequestXLanguageEnum struct {
	EN_US ShowDbObjectTemplateProgressRequestXLanguage
	ZH_CN ShowDbObjectTemplateProgressRequestXLanguage
}

func GetShowDbObjectTemplateProgressRequestXLanguageEnum() ShowDbObjectTemplateProgressRequestXLanguageEnum {
	return ShowDbObjectTemplateProgressRequestXLanguageEnum{
		EN_US: ShowDbObjectTemplateProgressRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowDbObjectTemplateProgressRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowDbObjectTemplateProgressRequestXLanguage) Value() string {
	return c.value
}

func (c ShowDbObjectTemplateProgressRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDbObjectTemplateProgressRequestXLanguage) UnmarshalJSON(b []byte) error {
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
