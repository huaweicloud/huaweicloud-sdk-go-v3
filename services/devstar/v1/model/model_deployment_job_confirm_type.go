package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DeploymentJobConfirmType struct {

	// 部署人工审核确认类型 stop终止部署 continue继续部署
	Type *DeploymentJobConfirmTypeType `json:"type,omitempty"`
}

func (o DeploymentJobConfirmType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentJobConfirmType struct{}"
	}

	return strings.Join([]string{"DeploymentJobConfirmType", string(data)}, " ")
}

type DeploymentJobConfirmTypeType struct {
	value string
}

type DeploymentJobConfirmTypeTypeEnum struct {
	STOP     DeploymentJobConfirmTypeType
	CONTINUE DeploymentJobConfirmTypeType
}

func GetDeploymentJobConfirmTypeTypeEnum() DeploymentJobConfirmTypeTypeEnum {
	return DeploymentJobConfirmTypeTypeEnum{
		STOP: DeploymentJobConfirmTypeType{
			value: "stop",
		},
		CONTINUE: DeploymentJobConfirmTypeType{
			value: "continue",
		},
	}
}

func (c DeploymentJobConfirmTypeType) Value() string {
	return c.value
}

func (c DeploymentJobConfirmTypeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeploymentJobConfirmTypeType) UnmarshalJSON(b []byte) error {
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
