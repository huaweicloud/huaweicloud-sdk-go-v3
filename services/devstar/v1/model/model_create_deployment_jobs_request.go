package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type CreateDeploymentJobsRequest struct {

	// 语言类型 中文:zh-cn 英文:en-us
	XLanguage *CreateDeploymentJobsRequestXLanguage `json:"X-Language,omitempty" xml:"X-Language"`

	// 应用id
	ApplicationId string `json:"application_id" xml:"application_id"`

	// 环境标识，从 [应用详情接口](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=DevStar&api=ShowApplication) 返回报文中的环境信息获取。
	EnvironmentTag string `json:"environment_tag" xml:"environment_tag"`

	Body *CreateDeploymentJobsParams `json:"body,omitempty" xml:"body"`
}

func (o CreateDeploymentJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeploymentJobsRequest struct{}"
	}

	return strings.Join([]string{"CreateDeploymentJobsRequest", string(data)}, " ")
}

type CreateDeploymentJobsRequestXLanguage struct {
	value string
}

type CreateDeploymentJobsRequestXLanguageEnum struct {
	ZH_CN CreateDeploymentJobsRequestXLanguage
	EN_US CreateDeploymentJobsRequestXLanguage
}

func GetCreateDeploymentJobsRequestXLanguageEnum() CreateDeploymentJobsRequestXLanguageEnum {
	return CreateDeploymentJobsRequestXLanguageEnum{
		ZH_CN: CreateDeploymentJobsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CreateDeploymentJobsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CreateDeploymentJobsRequestXLanguage) Value() string {
	return c.value
}

func (c CreateDeploymentJobsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDeploymentJobsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
