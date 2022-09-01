package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TemplateQuery struct {

	// 模板分类数组。
	Category *[]string `json:"category,omitempty" xml:"category"`

	// 搜索关键字，支持按名称和描述搜索，默认null。
	Keyword *string `json:"keyword,omitempty" xml:"keyword"`

	// 排序字段。
	SortKey *[]string `json:"sort_key,omitempty" xml:"sort_key"`

	// 指定排序使用升序还是降序。 - asc 升序 - desc 降序
	SortDir *[]string `json:"sort_dir,omitempty" xml:"sort_dir"`

	// 标签： - all：全部 - new：最新 - hot：热门 - recommend：推荐
	Label *TemplateQueryLabel `json:"label,omitempty" xml:"label"`

	// 是否查询用户自己创建的模板，默认查所有模板。
	MyTemplates *bool `json:"my_templates,omitempty" xml:"my_templates"`

	// 查所有模板时只处理上架的；查用户模板，需支持按状态查询，状态： - 0：审核中 - 1：上架 - 2：下架 不传表示查所有的（默认）
	Status *int32 `json:"status,omitempty" xml:"status"`

	// 模板状态数组。
	StatusArray *[]int32 `json:"status_array,omitempty" xml:"status_array"`

	// 是否查询有消息的模板，默认查所有模板。
	HasNotices *bool `json:"has_notices,omitempty" xml:"has_notices"`

	// 模板关联的云产品(产品短名)列表。
	Productshorts *[]string `json:"productshorts,omitempty" xml:"productshorts"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页的模板条数。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 模板关联的自定义标签列表。
	TagIds *[]string `json:"tag_ids,omitempty" xml:"tag_ids"`

	// 模板类型： - 0：doc - 1：code - 2：pipeline - 3：devops
	Types *[]int32 `json:"types,omitempty" xml:"types"`

	// 动、静态代码模板标识： - 0：动态模板codetemplate - 1：静态模板codesample
	IsStatic *int32 `json:"is_static,omitempty" xml:"is_static"`

	// 平台来源： - 0：codelabs - 1：devstar
	PlatformSource *[]int32 `json:"platform_source,omitempty" xml:"platform_source"`

	// 模板关联的标签名称列表。
	TagNames *[]string `json:"tag_names,omitempty" xml:"tag_names"`
}

func (o TemplateQuery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateQuery struct{}"
	}

	return strings.Join([]string{"TemplateQuery", string(data)}, " ")
}

type TemplateQueryLabel struct {
	value string
}

type TemplateQueryLabelEnum struct {
	ALL       TemplateQueryLabel
	NEW       TemplateQueryLabel
	HOT       TemplateQueryLabel
	RECOMMEND TemplateQueryLabel
}

func GetTemplateQueryLabelEnum() TemplateQueryLabelEnum {
	return TemplateQueryLabelEnum{
		ALL: TemplateQueryLabel{
			value: "all",
		},
		NEW: TemplateQueryLabel{
			value: "new",
		},
		HOT: TemplateQueryLabel{
			value: "hot",
		},
		RECOMMEND: TemplateQueryLabel{
			value: "recommend",
		},
	}
}

func (c TemplateQueryLabel) Value() string {
	return c.value
}

func (c TemplateQueryLabel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TemplateQueryLabel) UnmarshalJSON(b []byte) error {
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
