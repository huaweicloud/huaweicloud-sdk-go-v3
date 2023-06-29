package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SearchPublishInfoRequest Request Object
type SearchPublishInfoRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// dlm版本类型
	DlmType SearchPublishInfoRequestDlmType `json:"Dlm-Type"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// api编号
	ApiId string `json:"api_id"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`
}

func (o SearchPublishInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchPublishInfoRequest struct{}"
	}

	return strings.Join([]string{"SearchPublishInfoRequest", string(data)}, " ")
}

type SearchPublishInfoRequestDlmType struct {
	value string
}

type SearchPublishInfoRequestDlmTypeEnum struct {
	SHARED    SearchPublishInfoRequestDlmType
	EXCLUSIVE SearchPublishInfoRequestDlmType
}

func GetSearchPublishInfoRequestDlmTypeEnum() SearchPublishInfoRequestDlmTypeEnum {
	return SearchPublishInfoRequestDlmTypeEnum{
		SHARED: SearchPublishInfoRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: SearchPublishInfoRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c SearchPublishInfoRequestDlmType) Value() string {
	return c.value
}

func (c SearchPublishInfoRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchPublishInfoRequestDlmType) UnmarshalJSON(b []byte) error {
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
