package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateTaskV2RequestBody struct {

	// 部署任务id
	Id *string `json:"id,omitempty"`

	// 部署系统，deployTemplate：部署模板
	DeploySystem *UpdateTaskV2RequestBodyDeploySystem `json:"deploy_system,omitempty"`

	// 部署模板实例id
	TemplateId *string `json:"template_id,omitempty"`

	// 部署编排列表信息
	OperationList *[]DeployV2OperationsDo `json:"operation_list,omitempty"`
}

func (o UpdateTaskV2RequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskV2RequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTaskV2RequestBody", string(data)}, " ")
}

type UpdateTaskV2RequestBodyDeploySystem struct {
	value string
}

type UpdateTaskV2RequestBodyDeploySystemEnum struct {
	DEPLOY_TEMPLATE UpdateTaskV2RequestBodyDeploySystem
}

func GetUpdateTaskV2RequestBodyDeploySystemEnum() UpdateTaskV2RequestBodyDeploySystemEnum {
	return UpdateTaskV2RequestBodyDeploySystemEnum{
		DEPLOY_TEMPLATE: UpdateTaskV2RequestBodyDeploySystem{
			value: "deployTemplate",
		},
	}
}

func (c UpdateTaskV2RequestBodyDeploySystem) Value() string {
	return c.value
}

func (c UpdateTaskV2RequestBodyDeploySystem) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTaskV2RequestBodyDeploySystem) UnmarshalJSON(b []byte) error {
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
