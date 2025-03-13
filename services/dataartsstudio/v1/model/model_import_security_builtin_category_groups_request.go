package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ImportSecurityBuiltinCategoryGroupsRequest Request Object
type ImportSecurityBuiltinCategoryGroupsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 请求语言 * zh-cn 中文 * en-us 英文
	XLanguage ImportSecurityBuiltinCategoryGroupsRequestXLanguage `json:"X-Language"`

	Body *ImportBuiltinCategoryParam `json:"body,omitempty"`
}

func (o ImportSecurityBuiltinCategoryGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportSecurityBuiltinCategoryGroupsRequest struct{}"
	}

	return strings.Join([]string{"ImportSecurityBuiltinCategoryGroupsRequest", string(data)}, " ")
}

type ImportSecurityBuiltinCategoryGroupsRequestXLanguage struct {
	value string
}

type ImportSecurityBuiltinCategoryGroupsRequestXLanguageEnum struct {
	ZH_CN ImportSecurityBuiltinCategoryGroupsRequestXLanguage
	EN_US ImportSecurityBuiltinCategoryGroupsRequestXLanguage
}

func GetImportSecurityBuiltinCategoryGroupsRequestXLanguageEnum() ImportSecurityBuiltinCategoryGroupsRequestXLanguageEnum {
	return ImportSecurityBuiltinCategoryGroupsRequestXLanguageEnum{
		ZH_CN: ImportSecurityBuiltinCategoryGroupsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ImportSecurityBuiltinCategoryGroupsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ImportSecurityBuiltinCategoryGroupsRequestXLanguage) Value() string {
	return c.value
}

func (c ImportSecurityBuiltinCategoryGroupsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportSecurityBuiltinCategoryGroupsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
