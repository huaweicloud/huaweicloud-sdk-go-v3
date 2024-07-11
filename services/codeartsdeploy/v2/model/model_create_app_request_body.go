package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAppRequestBody 创建应用请求体
type CreateAppRequestBody struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 应用名称
	Name string `json:"name"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 分组id
	GroupId *string `json:"group_id,omitempty"`

	// 是否为草稿
	IsDraft bool `json:"is_draft"`

	// 创建类型，创建类型只有一个'template'，即根据模板创建
	CreateType CreateAppRequestBodyCreateType `json:"create_type"`

	// 自定义slave资源池id
	SlaveClusterId *string `json:"slave_cluster_id,omitempty"`

	Trigger *TaskTriggerVo `json:"trigger,omitempty"`

	// 部署任务列表信息
	ArrangeInfos *[]TaskV2RequestBody `json:"arrange_infos,omitempty"`
}

func (o CreateAppRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAppRequestBody", string(data)}, " ")
}

type CreateAppRequestBodyCreateType struct {
	value string
}

type CreateAppRequestBodyCreateTypeEnum struct {
	TEMPLATE CreateAppRequestBodyCreateType
}

func GetCreateAppRequestBodyCreateTypeEnum() CreateAppRequestBodyCreateTypeEnum {
	return CreateAppRequestBodyCreateTypeEnum{
		TEMPLATE: CreateAppRequestBodyCreateType{
			value: "template",
		},
	}
}

func (c CreateAppRequestBodyCreateType) Value() string {
	return c.value
}

func (c CreateAppRequestBodyCreateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAppRequestBodyCreateType) UnmarshalJSON(b []byte) error {
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
