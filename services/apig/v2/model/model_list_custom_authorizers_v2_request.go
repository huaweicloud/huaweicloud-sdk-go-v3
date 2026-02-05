package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCustomAuthorizersV2Request Request Object
type ListCustomAuthorizersV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量，条目数量小于等于0时，自动转换为20，条目数量大于500时，自动转换为500
	Limit *int32 `json:"limit,omitempty"`

	// 编号
	Id *string `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 自定义认证类型。 - FRONTEND：前端 - BACKEND：后端
	Type *ListCustomAuthorizersV2RequestType `json:"type,omitempty"`
}

func (o ListCustomAuthorizersV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomAuthorizersV2Request struct{}"
	}

	return strings.Join([]string{"ListCustomAuthorizersV2Request", string(data)}, " ")
}

type ListCustomAuthorizersV2RequestType struct {
	value string
}

type ListCustomAuthorizersV2RequestTypeEnum struct {
	FRONTEND ListCustomAuthorizersV2RequestType
	BACKEND  ListCustomAuthorizersV2RequestType
}

func GetListCustomAuthorizersV2RequestTypeEnum() ListCustomAuthorizersV2RequestTypeEnum {
	return ListCustomAuthorizersV2RequestTypeEnum{
		FRONTEND: ListCustomAuthorizersV2RequestType{
			value: "FRONTEND",
		},
		BACKEND: ListCustomAuthorizersV2RequestType{
			value: "BACKEND",
		},
	}
}

func (c ListCustomAuthorizersV2RequestType) Value() string {
	return c.value
}

func (c ListCustomAuthorizersV2RequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCustomAuthorizersV2RequestType) UnmarshalJSON(b []byte) error {
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
