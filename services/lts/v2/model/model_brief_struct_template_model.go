package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 结构化模板简略对象
type BriefStructTemplateModel struct {

	// 模板创建/更新时间
	CreateTime int64 `json:"create_time"`

	// 模板id
	Id string `json:"id"`

	// 模板名称
	TemplateName string `json:"template_name"`

	// 结构化类型，当前支持regex,json,split,nginx
	TemplateType BriefStructTemplateModelTemplateType `json:"template_type"`

	// 项目ID，获取方式请参见：获取账号ID、项目ID、日志组ID、日志流ID（https://support.huaweicloud.com/api-lts/lts_api_0006.html）。
	ProjectId string `json:"project_id"`
}

func (o BriefStructTemplateModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BriefStructTemplateModel struct{}"
	}

	return strings.Join([]string{"BriefStructTemplateModel", string(data)}, " ")
}

type BriefStructTemplateModelTemplateType struct {
	value string
}

type BriefStructTemplateModelTemplateTypeEnum struct {
	REGEX BriefStructTemplateModelTemplateType
	JSON  BriefStructTemplateModelTemplateType
	SPLIT BriefStructTemplateModelTemplateType
	NGINX BriefStructTemplateModelTemplateType
}

func GetBriefStructTemplateModelTemplateTypeEnum() BriefStructTemplateModelTemplateTypeEnum {
	return BriefStructTemplateModelTemplateTypeEnum{
		REGEX: BriefStructTemplateModelTemplateType{
			value: "regex",
		},
		JSON: BriefStructTemplateModelTemplateType{
			value: "json",
		},
		SPLIT: BriefStructTemplateModelTemplateType{
			value: "split",
		},
		NGINX: BriefStructTemplateModelTemplateType{
			value: "nginx",
		},
	}
}

func (c BriefStructTemplateModelTemplateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BriefStructTemplateModelTemplateType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
