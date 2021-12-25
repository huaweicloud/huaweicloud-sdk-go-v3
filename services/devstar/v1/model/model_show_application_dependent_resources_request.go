package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowApplicationDependentResourcesRequest struct {
	// 语言类型 中文:zh-cn 英文:en-us

	XLanguage *ShowApplicationDependentResourcesRequestXLanguage `json:"X-Language,omitempty"`
	// 应用id

	ApplicationId string `json:"application_id"`
}

func (o ShowApplicationDependentResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApplicationDependentResourcesRequest struct{}"
	}

	return strings.Join([]string{"ShowApplicationDependentResourcesRequest", string(data)}, " ")
}

type ShowApplicationDependentResourcesRequestXLanguage struct {
	value string
}

type ShowApplicationDependentResourcesRequestXLanguageEnum struct {
	ZH_CN ShowApplicationDependentResourcesRequestXLanguage
	EN_US ShowApplicationDependentResourcesRequestXLanguage
}

func GetShowApplicationDependentResourcesRequestXLanguageEnum() ShowApplicationDependentResourcesRequestXLanguageEnum {
	return ShowApplicationDependentResourcesRequestXLanguageEnum{
		ZH_CN: ShowApplicationDependentResourcesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowApplicationDependentResourcesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowApplicationDependentResourcesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApplicationDependentResourcesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
