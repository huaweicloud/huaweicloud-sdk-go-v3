package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportUserGroupsRequest Request Object
type ExportUserGroupsRequest struct {

	// 用于分页查询，每页返回的个数，取值范围0~50。
	Limit *int32 `json:"limit,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始，必须与limit同时使用。
	Offset *int32 `json:"offset,omitempty"`

	// 用来匹配用户组的搜索关键字。根据组名模糊查询。
	Keyword *string `json:"keyword,omitempty"`

	// 语言。 - zh_CN：中文 - en_US：英文
	Language *ExportUserGroupsRequestLanguage `json:"language,omitempty"`
}

func (o ExportUserGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUserGroupsRequest struct{}"
	}

	return strings.Join([]string{"ExportUserGroupsRequest", string(data)}, " ")
}

type ExportUserGroupsRequestLanguage struct {
	value string
}

type ExportUserGroupsRequestLanguageEnum struct {
	ZH_CN ExportUserGroupsRequestLanguage
	EN_US ExportUserGroupsRequestLanguage
}

func GetExportUserGroupsRequestLanguageEnum() ExportUserGroupsRequestLanguageEnum {
	return ExportUserGroupsRequestLanguageEnum{
		ZH_CN: ExportUserGroupsRequestLanguage{
			value: "zh_CN",
		},
		EN_US: ExportUserGroupsRequestLanguage{
			value: "en_US",
		},
	}
}

func (c ExportUserGroupsRequestLanguage) Value() string {
	return c.value
}

func (c ExportUserGroupsRequestLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportUserGroupsRequestLanguage) UnmarshalJSON(b []byte) error {
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
