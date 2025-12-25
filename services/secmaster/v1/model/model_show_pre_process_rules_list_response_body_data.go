package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ShowPreProcessRulesListResponseBodyData struct {

	// 映射id
	Id *string `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 映射id
	ProjectId *string `json:"project_id,omitempty"`

	// 映射id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 映射id
	MappingId *string `json:"mapping_id,omitempty"`

	// 映射id
	MapperId *string `json:"mapper_id,omitempty"`

	// 映射id
	MapperTypeId *string `json:"mapper_type_id,omitempty"`

	// 预处理选项: drop-丢弃
	Action *ShowPreProcessRulesListResponseBodyDataAction `json:"action,omitempty"`

	// 表达式
	Expression *string `json:"expression,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o ShowPreProcessRulesListResponseBodyData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPreProcessRulesListResponseBodyData struct{}"
	}

	return strings.Join([]string{"ShowPreProcessRulesListResponseBodyData", string(data)}, " ")
}

type ShowPreProcessRulesListResponseBodyDataAction struct {
	value string
}

type ShowPreProcessRulesListResponseBodyDataActionEnum struct {
	DROP ShowPreProcessRulesListResponseBodyDataAction
}

func GetShowPreProcessRulesListResponseBodyDataActionEnum() ShowPreProcessRulesListResponseBodyDataActionEnum {
	return ShowPreProcessRulesListResponseBodyDataActionEnum{
		DROP: ShowPreProcessRulesListResponseBodyDataAction{
			value: "drop",
		},
	}
}

func (c ShowPreProcessRulesListResponseBodyDataAction) Value() string {
	return c.value
}

func (c ShowPreProcessRulesListResponseBodyDataAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPreProcessRulesListResponseBodyDataAction) UnmarshalJSON(b []byte) error {
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
