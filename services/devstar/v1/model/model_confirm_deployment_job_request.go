package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ConfirmDeploymentJobRequest struct {

	// 语言类型 中文:zh-cn 英文:en-us
	XLanguage *ConfirmDeploymentJobRequestXLanguage `json:"X-Language,omitempty"`

	// 应用id
	ApplicationId string `json:"application_id"`

	// 环境标识
	EnvironmentTag string `json:"environment_tag"`

	Body *DeploymentJobConfirmType `json:"body,omitempty"`
}

func (o ConfirmDeploymentJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmDeploymentJobRequest struct{}"
	}

	return strings.Join([]string{"ConfirmDeploymentJobRequest", string(data)}, " ")
}

type ConfirmDeploymentJobRequestXLanguage struct {
	value string
}

type ConfirmDeploymentJobRequestXLanguageEnum struct {
	ZH_CN ConfirmDeploymentJobRequestXLanguage
	EN_US ConfirmDeploymentJobRequestXLanguage
}

func GetConfirmDeploymentJobRequestXLanguageEnum() ConfirmDeploymentJobRequestXLanguageEnum {
	return ConfirmDeploymentJobRequestXLanguageEnum{
		ZH_CN: ConfirmDeploymentJobRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ConfirmDeploymentJobRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ConfirmDeploymentJobRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfirmDeploymentJobRequestXLanguage) UnmarshalJSON(b []byte) error {
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
