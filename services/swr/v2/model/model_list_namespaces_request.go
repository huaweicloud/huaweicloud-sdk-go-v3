package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListNamespacesRequest Request Object
type ListNamespacesRequest struct {

	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json
	ContentType ListNamespacesRequestContentType `json:"Content-Type"`

	// 组织名称。小写字母开头，后面跟小写字母、数字、小数点、下划线或中划线（其中下划线最多允许连续两个，小数点、下划线、中划线不能直接相连），小写字母或数字结尾，1-64个字符。
	Namespace *string `json:"namespace,omitempty"`

	// 应填写namespace::{namespace}|mode::{mode}。其中{namespace}是组织名称，{mode}如果不设置，查看有权限的组织列表；设置为visible，查看可见的组织列表（部分组织：仓库有权限，组织没有权限）。
	Filter *string `json:"filter,omitempty"`
}

func (o ListNamespacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNamespacesRequest struct{}"
	}

	return strings.Join([]string{"ListNamespacesRequest", string(data)}, " ")
}

type ListNamespacesRequestContentType struct {
	value string
}

type ListNamespacesRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ListNamespacesRequestContentType
	APPLICATION_JSON             ListNamespacesRequestContentType
}

func GetListNamespacesRequestContentTypeEnum() ListNamespacesRequestContentTypeEnum {
	return ListNamespacesRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ListNamespacesRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ListNamespacesRequestContentType{
			value: "application/json",
		},
	}
}

func (c ListNamespacesRequestContentType) Value() string {
	return c.value
}

func (c ListNamespacesRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNamespacesRequestContentType) UnmarshalJSON(b []byte) error {
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
