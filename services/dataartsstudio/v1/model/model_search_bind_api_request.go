package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SearchBindApiRequest Request Object
type SearchBindApiRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// dlm版本类型
	DlmType *SearchBindApiRequestDlmType `json:"Dlm-Type,omitempty"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// app编号
	AppId string `json:"app_id"`

	// 查询起始坐标, 即跳过前X条数据。仅支持0或limit的整数倍，不满足则向下取整
	Offset *int32 `json:"offset,omitempty"`

	// 查询条数, 即查询Y条数据
	Limit *int32 `json:"limit,omitempty"`
}

func (o SearchBindApiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchBindApiRequest struct{}"
	}

	return strings.Join([]string{"SearchBindApiRequest", string(data)}, " ")
}

type SearchBindApiRequestDlmType struct {
	value string
}

type SearchBindApiRequestDlmTypeEnum struct {
	SHARED    SearchBindApiRequestDlmType
	EXCLUSIVE SearchBindApiRequestDlmType
}

func GetSearchBindApiRequestDlmTypeEnum() SearchBindApiRequestDlmTypeEnum {
	return SearchBindApiRequestDlmTypeEnum{
		SHARED: SearchBindApiRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: SearchBindApiRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c SearchBindApiRequestDlmType) Value() string {
	return c.value
}

func (c SearchBindApiRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchBindApiRequestDlmType) UnmarshalJSON(b []byte) error {
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
