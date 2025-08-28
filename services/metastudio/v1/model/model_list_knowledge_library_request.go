package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListKnowledgeLibraryRequest Request Object
type ListKnowledgeLibraryRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 角色语言  * CN：中文  * EN：英文
	Language *ListKnowledgeLibraryRequestLanguage `json:"language,omitempty"`

	// 知识库类型  * QUESTION_ANSWER：问答  * DOCUMENT：文档
	KnowledgeType *ListKnowledgeLibraryRequestKnowledgeType `json:"knowledge_type,omitempty"`

	// 知识库名称
	Name *string `json:"name,omitempty"`
}

func (o ListKnowledgeLibraryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKnowledgeLibraryRequest struct{}"
	}

	return strings.Join([]string{"ListKnowledgeLibraryRequest", string(data)}, " ")
}

type ListKnowledgeLibraryRequestLanguage struct {
	value string
}

type ListKnowledgeLibraryRequestLanguageEnum struct {
	CN ListKnowledgeLibraryRequestLanguage
	EN ListKnowledgeLibraryRequestLanguage
}

func GetListKnowledgeLibraryRequestLanguageEnum() ListKnowledgeLibraryRequestLanguageEnum {
	return ListKnowledgeLibraryRequestLanguageEnum{
		CN: ListKnowledgeLibraryRequestLanguage{
			value: "CN",
		},
		EN: ListKnowledgeLibraryRequestLanguage{
			value: "EN",
		},
	}
}

func (c ListKnowledgeLibraryRequestLanguage) Value() string {
	return c.value
}

func (c ListKnowledgeLibraryRequestLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListKnowledgeLibraryRequestLanguage) UnmarshalJSON(b []byte) error {
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

type ListKnowledgeLibraryRequestKnowledgeType struct {
	value string
}

type ListKnowledgeLibraryRequestKnowledgeTypeEnum struct {
	QUESTION_ANSWER ListKnowledgeLibraryRequestKnowledgeType
	DOCUMENT        ListKnowledgeLibraryRequestKnowledgeType
}

func GetListKnowledgeLibraryRequestKnowledgeTypeEnum() ListKnowledgeLibraryRequestKnowledgeTypeEnum {
	return ListKnowledgeLibraryRequestKnowledgeTypeEnum{
		QUESTION_ANSWER: ListKnowledgeLibraryRequestKnowledgeType{
			value: "QUESTION_ANSWER",
		},
		DOCUMENT: ListKnowledgeLibraryRequestKnowledgeType{
			value: "DOCUMENT",
		},
	}
}

func (c ListKnowledgeLibraryRequestKnowledgeType) Value() string {
	return c.value
}

func (c ListKnowledgeLibraryRequestKnowledgeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListKnowledgeLibraryRequestKnowledgeType) UnmarshalJSON(b []byte) error {
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
