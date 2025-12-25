package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCollectorModuleTemplateRequest Request Object
type ListCollectorModuleTemplateRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**: 解析类型 - FILTER 过滤 - INPUT 输入 - OUTPUT 输出  **约束限制** 不涉及 **取值范围**: - FILTER - INPUT - OUTPUT  **默认值** 不涉及
	ParserType *ListCollectorModuleTemplateRequestParserType `json:"parser_type,omitempty"`
}

func (o ListCollectorModuleTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectorModuleTemplateRequest struct{}"
	}

	return strings.Join([]string{"ListCollectorModuleTemplateRequest", string(data)}, " ")
}

type ListCollectorModuleTemplateRequestParserType struct {
	value string
}

type ListCollectorModuleTemplateRequestParserTypeEnum struct {
	FILTER ListCollectorModuleTemplateRequestParserType
	INPUT  ListCollectorModuleTemplateRequestParserType
	OUTPUT ListCollectorModuleTemplateRequestParserType
}

func GetListCollectorModuleTemplateRequestParserTypeEnum() ListCollectorModuleTemplateRequestParserTypeEnum {
	return ListCollectorModuleTemplateRequestParserTypeEnum{
		FILTER: ListCollectorModuleTemplateRequestParserType{
			value: "FILTER",
		},
		INPUT: ListCollectorModuleTemplateRequestParserType{
			value: "INPUT",
		},
		OUTPUT: ListCollectorModuleTemplateRequestParserType{
			value: "OUTPUT",
		},
	}
}

func (c ListCollectorModuleTemplateRequestParserType) Value() string {
	return c.value
}

func (c ListCollectorModuleTemplateRequestParserType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCollectorModuleTemplateRequestParserType) UnmarshalJSON(b []byte) error {
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
