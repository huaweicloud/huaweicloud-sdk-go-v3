package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListMessageRequest Request Object
type ListMessageRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 数据服务的版本类型，指定SHARED共享版或EXCLUSIVE专享版。
	DlmType *ListMessageRequestDlmType `json:"Dlm-Type,omitempty"`

	// 消息体的类型（格式），有Body体的情况下必选，没有Body体无需填写。如果请求消息体中含有中文字符，则需要通过charset=utf8指定中文字符集，例如取值为：application/json;charset=utf8。
	ContentType string `json:"Content-Type"`

	// 查询起始坐标, 即跳过前X条数据。仅支持0或limit的整数倍，不满足则向下取整。
	Offset *int32 `json:"offset,omitempty"`

	// 查询条数, 即查询Y条数据。
	Limit *int32 `json:"limit,omitempty"`

	// api名称。
	ApiName *string `json:"api_name,omitempty"`
}

func (o ListMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessageRequest struct{}"
	}

	return strings.Join([]string{"ListMessageRequest", string(data)}, " ")
}

type ListMessageRequestDlmType struct {
	value string
}

type ListMessageRequestDlmTypeEnum struct {
	SHARED    ListMessageRequestDlmType
	EXCLUSIVE ListMessageRequestDlmType
}

func GetListMessageRequestDlmTypeEnum() ListMessageRequestDlmTypeEnum {
	return ListMessageRequestDlmTypeEnum{
		SHARED: ListMessageRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ListMessageRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ListMessageRequestDlmType) Value() string {
	return c.value
}

func (c ListMessageRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMessageRequestDlmType) UnmarshalJSON(b []byte) error {
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
