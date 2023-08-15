package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDeploymentJobsRequest Request Object
type ShowDeploymentJobsRequest struct {

	// 语言类型 中文:zh-cn 英文:en-us
	XLanguage *ShowDeploymentJobsRequestXLanguage `json:"X-Language,omitempty"`

	// 应用id
	ApplicationId string `json:"application_id"`

	// 环境标识，从 [应用详情接口](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=DevStar&api=ShowApplication) 返回报文中的环境信息获取。
	EnvironmentTag string `json:"environment_tag"`
}

func (o ShowDeploymentJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeploymentJobsRequest struct{}"
	}

	return strings.Join([]string{"ShowDeploymentJobsRequest", string(data)}, " ")
}

type ShowDeploymentJobsRequestXLanguage struct {
	value string
}

type ShowDeploymentJobsRequestXLanguageEnum struct {
	ZH_CN ShowDeploymentJobsRequestXLanguage
	EN_US ShowDeploymentJobsRequestXLanguage
}

func GetShowDeploymentJobsRequestXLanguageEnum() ShowDeploymentJobsRequestXLanguageEnum {
	return ShowDeploymentJobsRequestXLanguageEnum{
		ZH_CN: ShowDeploymentJobsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowDeploymentJobsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowDeploymentJobsRequestXLanguage) Value() string {
	return c.value
}

func (c ShowDeploymentJobsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDeploymentJobsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
