package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ShowRepositoryStatisticalDataV2Request struct {
	// 语言类型 中文:zh-cn 英文:en-us

	XLanguage *ShowRepositoryStatisticalDataV2RequestXLanguage `json:"X-Language,omitempty"`
	// 代码仓库id

	RepositoryId string `json:"repository_id"`
}

func (o ShowRepositoryStatisticalDataV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryStatisticalDataV2Request struct{}"
	}

	return strings.Join([]string{"ShowRepositoryStatisticalDataV2Request", string(data)}, " ")
}

type ShowRepositoryStatisticalDataV2RequestXLanguage struct {
	value string
}

type ShowRepositoryStatisticalDataV2RequestXLanguageEnum struct {
	ZH_CN ShowRepositoryStatisticalDataV2RequestXLanguage
	EN_US ShowRepositoryStatisticalDataV2RequestXLanguage
}

func GetShowRepositoryStatisticalDataV2RequestXLanguageEnum() ShowRepositoryStatisticalDataV2RequestXLanguageEnum {
	return ShowRepositoryStatisticalDataV2RequestXLanguageEnum{
		ZH_CN: ShowRepositoryStatisticalDataV2RequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowRepositoryStatisticalDataV2RequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowRepositoryStatisticalDataV2RequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRepositoryStatisticalDataV2RequestXLanguage) UnmarshalJSON(b []byte) error {
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
