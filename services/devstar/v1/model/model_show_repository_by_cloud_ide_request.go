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

	// 仓库id。
	RepositoryId string `json:"repository_id"`

	// 仓库下载地址。
	RepositorySshUrl string `json:"repository_ssh_url"`

	// 区域ID，目前仅支持北京四：cn-north-4及北京一：cn-north-1。
	RegionId *string `json:"region_id,omitempty"`

	// 工作空间名称前缀，仅在is_open_last为false时生效，由用户自定义，支持大小写字母、中文、_、-，长度1-256。
	SpacePrefix *string `json:"space_prefix,omitempty"`

	// 是否打开上一次的工作空间，true表示打开上一次工作空间，如果没有上一次工作空间会返回空，false代表打开一个全新的工作空间。
	IsOpenLast *bool `json:"is_open_last,omitempty"`

	// 是否创建 CloudIDE 免费实例链接，true表示创建一个 CloudIDE 免费实例链接，false表示创建一个 CloudIDE 收费实例链接。
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
