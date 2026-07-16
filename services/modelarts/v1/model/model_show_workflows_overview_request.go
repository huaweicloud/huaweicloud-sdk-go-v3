package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowWorkflowsOverviewRequest Request Object
type ShowWorkflowsOverviewRequest struct {

	// 工作空间ID。[获取方法请参见[查询工作空间列表](ListWorkspace.xml)。](tag:hc)未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 过滤方式。可选值如下： - equal表示精确匹配。 - contain表示模糊匹配。  具体过滤的字段，由各个接口额外定义参数。例如Workflow支持按照名称（name）进行过滤，则相应的过滤字段为name。name=workflow&search_type=contain表示查询名称中含有Workflow字样的所有工作流。
	SearchType *ShowWorkflowsOverviewRequestSearchType `json:"search_type,omitempty"`

	// 工作流名称。填写1-64位，仅包含英文、数字、下划线（_）和中划线（-），并且以英文开头的名称。
	Name *string `json:"name,omitempty"`

	// 工作流描述信息。
	Description *string `json:"description,omitempty"`
}

func (o ShowWorkflowsOverviewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkflowsOverviewRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkflowsOverviewRequest", string(data)}, " ")
}

type ShowWorkflowsOverviewRequestSearchType struct {
	value string
}

type ShowWorkflowsOverviewRequestSearchTypeEnum struct {
	CONTAIN ShowWorkflowsOverviewRequestSearchType
	EQUAL   ShowWorkflowsOverviewRequestSearchType
}

func GetShowWorkflowsOverviewRequestSearchTypeEnum() ShowWorkflowsOverviewRequestSearchTypeEnum {
	return ShowWorkflowsOverviewRequestSearchTypeEnum{
		CONTAIN: ShowWorkflowsOverviewRequestSearchType{
			value: "contain",
		},
		EQUAL: ShowWorkflowsOverviewRequestSearchType{
			value: "equal",
		},
	}
}

func (c ShowWorkflowsOverviewRequestSearchType) Value() string {
	return c.value
}

func (c ShowWorkflowsOverviewRequestSearchType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowWorkflowsOverviewRequestSearchType) UnmarshalJSON(b []byte) error {
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
