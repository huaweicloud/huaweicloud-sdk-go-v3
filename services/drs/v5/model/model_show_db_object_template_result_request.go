package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDbObjectTemplateResultRequest Request Object
type ShowDbObjectTemplateResultRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ShowDbObjectTemplateResultRequestXLanguage `json:"X-Language,omitempty"`

	// 导入的结果类型。取值： - detail：获取最新导入的文件与校验结果，上传后的文件若存在错误，会同时将错误原因标记在文件内。 - synchronized：获取已同步的（已下发的）对象文件结果。 - change: 获取新增和删除的对象结果（当再编辑时使用）
	Type ShowDbObjectTemplateResultRequestType `json:"type"`

	// 默认为空。当进行列加工导入时，查询列加工导入进度，取值column。
	FileExportObjectLevel *string `json:"file_export_object_level,omitempty"`
}

func (o ShowDbObjectTemplateResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDbObjectTemplateResultRequest struct{}"
	}

	return strings.Join([]string{"ShowDbObjectTemplateResultRequest", string(data)}, " ")
}

type ShowDbObjectTemplateResultRequestXLanguage struct {
	value string
}

type ShowDbObjectTemplateResultRequestXLanguageEnum struct {
	EN_US ShowDbObjectTemplateResultRequestXLanguage
	ZH_CN ShowDbObjectTemplateResultRequestXLanguage
}

func GetShowDbObjectTemplateResultRequestXLanguageEnum() ShowDbObjectTemplateResultRequestXLanguageEnum {
	return ShowDbObjectTemplateResultRequestXLanguageEnum{
		EN_US: ShowDbObjectTemplateResultRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowDbObjectTemplateResultRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowDbObjectTemplateResultRequestXLanguage) Value() string {
	return c.value
}

func (c ShowDbObjectTemplateResultRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDbObjectTemplateResultRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ShowDbObjectTemplateResultRequestType struct {
	value string
}

type ShowDbObjectTemplateResultRequestTypeEnum struct {
	DETAIL       ShowDbObjectTemplateResultRequestType
	SYNCHRONIZED ShowDbObjectTemplateResultRequestType
	CHANGE       ShowDbObjectTemplateResultRequestType
}

func GetShowDbObjectTemplateResultRequestTypeEnum() ShowDbObjectTemplateResultRequestTypeEnum {
	return ShowDbObjectTemplateResultRequestTypeEnum{
		DETAIL: ShowDbObjectTemplateResultRequestType{
			value: "detail",
		},
		SYNCHRONIZED: ShowDbObjectTemplateResultRequestType{
			value: "synchronized",
		},
		CHANGE: ShowDbObjectTemplateResultRequestType{
			value: "change",
		},
	}
}

func (c ShowDbObjectTemplateResultRequestType) Value() string {
	return c.value
}

func (c ShowDbObjectTemplateResultRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDbObjectTemplateResultRequestType) UnmarshalJSON(b []byte) error {
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
