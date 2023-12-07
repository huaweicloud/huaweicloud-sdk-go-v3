package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListApisRequest Request Object
type ListApisRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// dlm版本类型
	DlmType *ListApisRequestDlmType `json:"Dlm-Type,omitempty"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// 是否返回专享版API的发布信息
	XReturnPublishMessages *string `json:"x-return-publish-messages,omitempty"`

	// 查询起始坐标, 即跳过前X条数据。仅支持0或limit的整数倍，不满足则向下取整
	Offset *int32 `json:"offset,omitempty"`

	// 查询条数, 即查询Y条数据
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListApisRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApisRequest struct{}"
	}

	return strings.Join([]string{"ListApisRequest", string(data)}, " ")
}

type ListApisRequestDlmType struct {
	value string
}

type ListApisRequestDlmTypeEnum struct {
	SHARED    ListApisRequestDlmType
	EXCLUSIVE ListApisRequestDlmType
}

func GetListApisRequestDlmTypeEnum() ListApisRequestDlmTypeEnum {
	return ListApisRequestDlmTypeEnum{
		SHARED: ListApisRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ListApisRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ListApisRequestDlmType) Value() string {
	return c.value
}

func (c ListApisRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApisRequestDlmType) UnmarshalJSON(b []byte) error {
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
