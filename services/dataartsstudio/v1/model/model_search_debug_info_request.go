package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SearchDebugInfoRequest Request Object
type SearchDebugInfoRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// dlm版本类型
	DlmType SearchDebugInfoRequestDlmType `json:"Dlm-Type"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// api编号
	ApiId string `json:"api_id"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`
}

func (o SearchDebugInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchDebugInfoRequest struct{}"
	}

	return strings.Join([]string{"SearchDebugInfoRequest", string(data)}, " ")
}

type SearchDebugInfoRequestDlmType struct {
	value string
}

type SearchDebugInfoRequestDlmTypeEnum struct {
	SHARED    SearchDebugInfoRequestDlmType
	EXCLUSIVE SearchDebugInfoRequestDlmType
}

func GetSearchDebugInfoRequestDlmTypeEnum() SearchDebugInfoRequestDlmTypeEnum {
	return SearchDebugInfoRequestDlmTypeEnum{
		SHARED: SearchDebugInfoRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: SearchDebugInfoRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c SearchDebugInfoRequestDlmType) Value() string {
	return c.value
}

func (c SearchDebugInfoRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchDebugInfoRequestDlmType) UnmarshalJSON(b []byte) error {
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
