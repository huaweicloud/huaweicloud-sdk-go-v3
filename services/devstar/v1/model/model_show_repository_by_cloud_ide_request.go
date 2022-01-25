package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowRepositoryByCloudIdeRequest struct {
	// 语言类型 中文:zh-cn 英文:en-us

	XLanguage *ShowRepositoryByCloudIdeRequestXLanguage `json:"X-Language,omitempty"`
	// 仓库id

	RepositoryId string `json:"repository_id"`
	// 仓库下载地址

	RepositorySshUrl string `json:"repository_ssh_url"`
	// 区域ID

	RegionId *string `json:"region_id,omitempty"`
	// 工作空间名称前缀

	SpacePrefix *string `json:"space_prefix,omitempty"`
	// 是否打开上一次的工作空间

	IsOpenLast *bool `json:"is_open_last,omitempty"`
	// 是否创建免费实例链接

	IsFree *bool `json:"is_free,omitempty"`
}

func (o ShowRepositoryByCloudIdeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryByCloudIdeRequest struct{}"
	}

	return strings.Join([]string{"ShowRepositoryByCloudIdeRequest", string(data)}, " ")
}

type ShowRepositoryByCloudIdeRequestXLanguage struct {
	value string
}

type ShowRepositoryByCloudIdeRequestXLanguageEnum struct {
	ZH_CN ShowRepositoryByCloudIdeRequestXLanguage
	EN_US ShowRepositoryByCloudIdeRequestXLanguage
}

func GetShowRepositoryByCloudIdeRequestXLanguageEnum() ShowRepositoryByCloudIdeRequestXLanguageEnum {
	return ShowRepositoryByCloudIdeRequestXLanguageEnum{
		ZH_CN: ShowRepositoryByCloudIdeRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowRepositoryByCloudIdeRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowRepositoryByCloudIdeRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRepositoryByCloudIdeRequestXLanguage) UnmarshalJSON(b []byte) error {
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
