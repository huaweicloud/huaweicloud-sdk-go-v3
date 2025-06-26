package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowInstanceArtifactAdditionRequest Request Object
type ShowInstanceArtifactAdditionRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 命名空间名称
	NamespaceName string `json:"namespace_name"`

	// 仓库名称
	RepositoryName string `json:"repository_name"`

	// 制品摘要
	Reference string `json:"reference"`

	// 制品附加信息
	Addition ShowInstanceArtifactAdditionRequestAddition `json:"addition"`
}

func (o ShowInstanceArtifactAdditionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceArtifactAdditionRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceArtifactAdditionRequest", string(data)}, " ")
}

type ShowInstanceArtifactAdditionRequestAddition struct {
	value string
}

type ShowInstanceArtifactAdditionRequestAdditionEnum struct {
	BUILD_HISTORY ShowInstanceArtifactAdditionRequestAddition
}

func GetShowInstanceArtifactAdditionRequestAdditionEnum() ShowInstanceArtifactAdditionRequestAdditionEnum {
	return ShowInstanceArtifactAdditionRequestAdditionEnum{
		BUILD_HISTORY: ShowInstanceArtifactAdditionRequestAddition{
			value: "build_history",
		},
	}
}

func (c ShowInstanceArtifactAdditionRequestAddition) Value() string {
	return c.value
}

func (c ShowInstanceArtifactAdditionRequestAddition) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceArtifactAdditionRequestAddition) UnmarshalJSON(b []byte) error {
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
