package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ExportUserGroupsNewReq struct {

	// 用户组名模糊查询关键字。
	Name *string `json:"name,omitempty"`

	// 用户组ID列表。有传则与name过滤结果取交集。
	GroupIds *[]string `json:"group_ids,omitempty"`

	// 语言，用于Excel标题国际化。 * zh_CN： 中文 * en_US： 英文
	Language *ExportUserGroupsNewReqLanguage `json:"language,omitempty"`
}

func (o ExportUserGroupsNewReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUserGroupsNewReq struct{}"
	}

	return strings.Join([]string{"ExportUserGroupsNewReq", string(data)}, " ")
}

type ExportUserGroupsNewReqLanguage struct {
	value string
}

type ExportUserGroupsNewReqLanguageEnum struct {
	ZH_CN ExportUserGroupsNewReqLanguage
	EN_US ExportUserGroupsNewReqLanguage
}

func GetExportUserGroupsNewReqLanguageEnum() ExportUserGroupsNewReqLanguageEnum {
	return ExportUserGroupsNewReqLanguageEnum{
		ZH_CN: ExportUserGroupsNewReqLanguage{
			value: "zh_CN",
		},
		EN_US: ExportUserGroupsNewReqLanguage{
			value: "en_US",
		},
	}
}

func (c ExportUserGroupsNewReqLanguage) Value() string {
	return c.value
}

func (c ExportUserGroupsNewReqLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportUserGroupsNewReqLanguage) UnmarshalJSON(b []byte) error {
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
