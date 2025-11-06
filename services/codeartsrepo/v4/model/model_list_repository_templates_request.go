package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListRepositoryTemplatesRequest Request Object
type ListRepositoryTemplatesRequest struct {

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 应用类型。 **约束限制：** 不涉及 **取值范围：** - Android。 - Console。 - HarmonyOS。 - OTHERS。 - REST API。 - ServiceStage。 - Web网站。 - 图形用户界面。 **默认取值：** 不涉及
	Platform *string `json:"platform,omitempty"`

	// **参数解释：** 是否支持自动创建流水线。 **约束限制：** 不涉及 **取值范围：** - SupportPipeline，支持自动创建流水线。 - UnsupportedPipeline 不支持自动创建流水线。 **默认取值：** 不涉及
	Pipeline *ListRepositoryTemplatesRequestPipeline `json:"pipeline,omitempty"`

	// **参数解释：** 自动创建流水线。 **约束限制：** 不涉及 **取值范围：** - SYSTEM,USER，个人和官方模板。 - SYSTEM，官方模板。 - USER，个人模板。 **默认取值：** 不涉及
	Type ListRepositoryTemplatesRequestType `json:"type"`

	// **参数解释：** 查询关键字，按模板仓标题搜索。 **约束限制：** 不涉及 **取值范围：** 字符串长度0至50。 **默认取值：** 不涉及
	Search *string `json:"search,omitempty"`

	// **参数解释：** 参赛类型。 **约束限制：** 不涉及 **取值范围：** - AI。 - 大数据。 - 小程序。 - 微服务。 **默认取值：** 不涉及
	EnterType *string `json:"enter_type,omitempty"`

	// **参数解释：** 按照模板仓的创建时间进行排序。 **约束限制：** 不涉及 **取值范围：** - up，升序。 - down，降序。 **默认取值：** 不涉及
	DateOrder *ListRepositoryTemplatesRequestDateOrder `json:"date_order,omitempty"`

	// **参数解释：** 编程语言。 **约束限制：** 不涉及 **取值范围：** - ArkTS。 - ASP.NET。 - C。 - C#。 - C++。 - Go。 - Groovy。 - HTML。 - Java。 - NodeJS。 - OTHERS。 - PHP。 - Python。 **默认取值：** 不涉及
	Language *string `json:"language,omitempty"`

	// **参数解释：** 项目的32位uuid，项目唯一标识，通过[[查询项目列表](https://support.huaweicloud.com/api-projectman/ListProjectsV4.html)](tag:hws)[[查询项目列表](https://support.huaweicloud.com/intl/en-us/api-projectman/ListProjectsV4.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 当该参数不为空时，仅返回当前项目下符合搜索条件的模板仓 **取值范围：** 字符串长度32。 **取值范围：** 不涉及
	ProjectId *string `json:"project_id,omitempty"`
}

func (o ListRepositoryTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListRepositoryTemplatesRequest", string(data)}, " ")
}

type ListRepositoryTemplatesRequestPipeline struct {
	value string
}

type ListRepositoryTemplatesRequestPipelineEnum struct {
	SUPPORT_PIPELINE     ListRepositoryTemplatesRequestPipeline
	UNSUPPORTED_PIPELINE ListRepositoryTemplatesRequestPipeline
}

func GetListRepositoryTemplatesRequestPipelineEnum() ListRepositoryTemplatesRequestPipelineEnum {
	return ListRepositoryTemplatesRequestPipelineEnum{
		SUPPORT_PIPELINE: ListRepositoryTemplatesRequestPipeline{
			value: "SupportPipeline",
		},
		UNSUPPORTED_PIPELINE: ListRepositoryTemplatesRequestPipeline{
			value: "UnsupportedPipeline",
		},
	}
}

func (c ListRepositoryTemplatesRequestPipeline) Value() string {
	return c.value
}

func (c ListRepositoryTemplatesRequestPipeline) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRepositoryTemplatesRequestPipeline) UnmarshalJSON(b []byte) error {
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

type ListRepositoryTemplatesRequestType struct {
	value string
}

type ListRepositoryTemplatesRequestTypeEnum struct {
	SYSTEMUSER ListRepositoryTemplatesRequestType
	SYSTEM     ListRepositoryTemplatesRequestType
	USER       ListRepositoryTemplatesRequestType
}

func GetListRepositoryTemplatesRequestTypeEnum() ListRepositoryTemplatesRequestTypeEnum {
	return ListRepositoryTemplatesRequestTypeEnum{
		SYSTEMUSER: ListRepositoryTemplatesRequestType{
			value: "SYSTEM,USER",
		},
		SYSTEM: ListRepositoryTemplatesRequestType{
			value: "SYSTEM",
		},
		USER: ListRepositoryTemplatesRequestType{
			value: "USER",
		},
	}
}

func (c ListRepositoryTemplatesRequestType) Value() string {
	return c.value
}

func (c ListRepositoryTemplatesRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRepositoryTemplatesRequestType) UnmarshalJSON(b []byte) error {
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

type ListRepositoryTemplatesRequestDateOrder struct {
	value string
}

type ListRepositoryTemplatesRequestDateOrderEnum struct {
	UP   ListRepositoryTemplatesRequestDateOrder
	DOWN ListRepositoryTemplatesRequestDateOrder
}

func GetListRepositoryTemplatesRequestDateOrderEnum() ListRepositoryTemplatesRequestDateOrderEnum {
	return ListRepositoryTemplatesRequestDateOrderEnum{
		UP: ListRepositoryTemplatesRequestDateOrder{
			value: "up",
		},
		DOWN: ListRepositoryTemplatesRequestDateOrder{
			value: "down",
		},
	}
}

func (c ListRepositoryTemplatesRequestDateOrder) Value() string {
	return c.value
}

func (c ListRepositoryTemplatesRequestDateOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRepositoryTemplatesRequestDateOrder) UnmarshalJSON(b []byte) error {
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
