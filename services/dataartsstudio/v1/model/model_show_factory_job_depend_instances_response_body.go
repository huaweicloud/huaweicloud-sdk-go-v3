package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ShowFactoryJobDependInstancesResponseBody struct {

	// 依赖的作业名称。
	JobName *string `json:"job_name,omitempty"`

	// 依赖的作业所在目录路径。作业在根目录下返回\"/\"。
	JobPath *string `json:"job_path,omitempty"`

	// 当前作业与查询目标作业的依赖关系方向。 取值范围： - parent：当前作业是查询目标作业的上游作业。 - child：当前作业是查询目标作业的下游作业。
	DependLayer *ShowFactoryJobDependInstancesResponseBodyDependLayer `json:"depend_layer,omitempty"`

	// 依赖的作业所在的工作空间名称。
	WorkspaceName *string `json:"workspace_name,omitempty"`

	// 作业责任人。创建作业时指定的作业负责人。
	Owner *string `json:"owner,omitempty"`
}

func (o ShowFactoryJobDependInstancesResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFactoryJobDependInstancesResponseBody struct{}"
	}

	return strings.Join([]string{"ShowFactoryJobDependInstancesResponseBody", string(data)}, " ")
}

type ShowFactoryJobDependInstancesResponseBodyDependLayer struct {
	value string
}

type ShowFactoryJobDependInstancesResponseBodyDependLayerEnum struct {
	PARENT ShowFactoryJobDependInstancesResponseBodyDependLayer
	CHILD  ShowFactoryJobDependInstancesResponseBodyDependLayer
}

func GetShowFactoryJobDependInstancesResponseBodyDependLayerEnum() ShowFactoryJobDependInstancesResponseBodyDependLayerEnum {
	return ShowFactoryJobDependInstancesResponseBodyDependLayerEnum{
		PARENT: ShowFactoryJobDependInstancesResponseBodyDependLayer{
			value: "parent",
		},
		CHILD: ShowFactoryJobDependInstancesResponseBodyDependLayer{
			value: "child",
		},
	}
}

func (c ShowFactoryJobDependInstancesResponseBodyDependLayer) Value() string {
	return c.value
}

func (c ShowFactoryJobDependInstancesResponseBodyDependLayer) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFactoryJobDependInstancesResponseBodyDependLayer) UnmarshalJSON(b []byte) error {
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
