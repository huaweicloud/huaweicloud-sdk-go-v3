package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowApplicationReleaseRepositoriesRequest struct {

	// 语言类型 中文:zh-cn 英文:en-us
	XLanguage *ShowApplicationReleaseRepositoriesRequestXLanguage `json:"X-Language,omitempty"`

	// 应用id
	ApplicationId string `json:"application_id"`

	// 父id,仅在仓库类型为ReleaseMan需要
	ParentId *string `json:"parent_id,omitempty"`

	// 搜索关键字,支持按名称搜索,默认null
	Keyword *string `json:"keyword,omitempty"`

	// 每页显示的条目数量,默认10
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询,默认0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowApplicationReleaseRepositoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApplicationReleaseRepositoriesRequest struct{}"
	}

	return strings.Join([]string{"ShowApplicationReleaseRepositoriesRequest", string(data)}, " ")
}

type ShowApplicationReleaseRepositoriesRequestXLanguage struct {
	value string
}

type ShowApplicationReleaseRepositoriesRequestXLanguageEnum struct {
	ZH_CN ShowApplicationReleaseRepositoriesRequestXLanguage
	EN_US ShowApplicationReleaseRepositoriesRequestXLanguage
}

func GetShowApplicationReleaseRepositoriesRequestXLanguageEnum() ShowApplicationReleaseRepositoriesRequestXLanguageEnum {
	return ShowApplicationReleaseRepositoriesRequestXLanguageEnum{
		ZH_CN: ShowApplicationReleaseRepositoriesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowApplicationReleaseRepositoriesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowApplicationReleaseRepositoriesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApplicationReleaseRepositoriesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
