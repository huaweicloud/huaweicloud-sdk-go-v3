package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SearchIdByPathRequest Request Object
type SearchIdByPathRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// dlm版本类型
	DlmType SearchIdByPathRequestDlmType `json:"Dlm-Type"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// 路径
	Path string `json:"path"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`
}

func (o SearchIdByPathRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchIdByPathRequest struct{}"
	}

	return strings.Join([]string{"SearchIdByPathRequest", string(data)}, " ")
}

type SearchIdByPathRequestDlmType struct {
	value string
}

type SearchIdByPathRequestDlmTypeEnum struct {
	SHARED    SearchIdByPathRequestDlmType
	EXCLUSIVE SearchIdByPathRequestDlmType
}

func GetSearchIdByPathRequestDlmTypeEnum() SearchIdByPathRequestDlmTypeEnum {
	return SearchIdByPathRequestDlmTypeEnum{
		SHARED: SearchIdByPathRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: SearchIdByPathRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c SearchIdByPathRequestDlmType) Value() string {
	return c.value
}

func (c SearchIdByPathRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchIdByPathRequestDlmType) UnmarshalJSON(b []byte) error {
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
