package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDirectoriesRequest Request Object
type ListDirectoriesRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**: 目录类型 - TABLE 表 - PIPE 管道 - RETRIEVE_SCRIPT 检索脚本 - ANALYSIS_SCRIPT 分析脚本 - DATA_TRANSFORMATION 数据加工 - ALERT_RULE 告警规则  **约束限制** 不涉及 **取值范围**: - TABLE - PIPE - RETRIEVE_SCRIPT - ANALYSIS_SCRIPT - DATA_TRANSFORMATION - ALERT_RULE  **默认值** 不涉及
	Category ListDirectoriesRequestCategory `json:"category"`
}

func (o ListDirectoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDirectoriesRequest struct{}"
	}

	return strings.Join([]string{"ListDirectoriesRequest", string(data)}, " ")
}

type ListDirectoriesRequestCategory struct {
	value string
}

type ListDirectoriesRequestCategoryEnum struct {
	TABLE               ListDirectoriesRequestCategory
	PIPE                ListDirectoriesRequestCategory
	RETRIEVE_SCRIPT     ListDirectoriesRequestCategory
	ANALYSIS_SCRIPT     ListDirectoriesRequestCategory
	DATA_TRANSFORMATION ListDirectoriesRequestCategory
	ALERT_RULE          ListDirectoriesRequestCategory
}

func GetListDirectoriesRequestCategoryEnum() ListDirectoriesRequestCategoryEnum {
	return ListDirectoriesRequestCategoryEnum{
		TABLE: ListDirectoriesRequestCategory{
			value: "TABLE",
		},
		PIPE: ListDirectoriesRequestCategory{
			value: "PIPE",
		},
		RETRIEVE_SCRIPT: ListDirectoriesRequestCategory{
			value: "RETRIEVE_SCRIPT",
		},
		ANALYSIS_SCRIPT: ListDirectoriesRequestCategory{
			value: "ANALYSIS_SCRIPT",
		},
		DATA_TRANSFORMATION: ListDirectoriesRequestCategory{
			value: "DATA_TRANSFORMATION",
		},
		ALERT_RULE: ListDirectoriesRequestCategory{
			value: "ALERT_RULE",
		},
	}
}

func (c ListDirectoriesRequestCategory) Value() string {
	return c.value
}

func (c ListDirectoriesRequestCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDirectoriesRequestCategory) UnmarshalJSON(b []byte) error {
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
