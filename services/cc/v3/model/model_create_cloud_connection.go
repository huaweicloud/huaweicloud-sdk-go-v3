package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 创建云连接实例的详细信息。
type CreateCloudConnection struct {
	// 云连接实例的名字。

	Name string `json:"name"`
	// 云连接实例的描述。

	Description *string `json:"description,omitempty"`
	// 云连接实例所属的企业项目ID。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 云连接使用场景。|- VPC：虚拟私有云。 ER：虚拟路由器。

	UsedScene *CreateCloudConnectionUsedScene `json:"used_scene,omitempty"`
}

func (o CreateCloudConnection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudConnection struct{}"
	}

	return strings.Join([]string{"CreateCloudConnection", string(data)}, " ")
}

type CreateCloudConnectionUsedScene struct {
	value string
}

type CreateCloudConnectionUsedSceneEnum struct {
	ER  CreateCloudConnectionUsedScene
	VPC CreateCloudConnectionUsedScene
}

func GetCreateCloudConnectionUsedSceneEnum() CreateCloudConnectionUsedSceneEnum {
	return CreateCloudConnectionUsedSceneEnum{
		ER: CreateCloudConnectionUsedScene{
			value: "er",
		},
		VPC: CreateCloudConnectionUsedScene{
			value: "vpc",
		},
	}
}

func (c CreateCloudConnectionUsedScene) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCloudConnectionUsedScene) UnmarshalJSON(b []byte) error {
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
