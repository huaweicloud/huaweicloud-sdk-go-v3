package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCesInstanceRequestBody 查询medialive维度(medialive_mpc、medialive_cdn、medialive_package、medialive_connect、medialive_tailor) medialive_mpc维度可以查询实例下的pipeline实例及pipeline下的音频实例
type ListCesInstanceRequestBody struct {

	// 命名空间
	Namespace ListCesInstanceRequestBodyNamespace `json:"namespace"`

	// 查询信息
	Query []ListCesInstanceRequestBodyQuery `json:"query"`

	// 查询开始偏移
	Start *int32 `json:"start,omitempty"`

	// 查询限制
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListCesInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCesInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"ListCesInstanceRequestBody", string(data)}, " ")
}

type ListCesInstanceRequestBodyNamespace struct {
	value string
}

type ListCesInstanceRequestBodyNamespaceEnum struct {
	SYS_LIVE ListCesInstanceRequestBodyNamespace
}

func GetListCesInstanceRequestBodyNamespaceEnum() ListCesInstanceRequestBodyNamespaceEnum {
	return ListCesInstanceRequestBodyNamespaceEnum{
		SYS_LIVE: ListCesInstanceRequestBodyNamespace{
			value: "SYS.LIVE",
		},
	}
}

func (c ListCesInstanceRequestBodyNamespace) Value() string {
	return c.value
}

func (c ListCesInstanceRequestBodyNamespace) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCesInstanceRequestBodyNamespace) UnmarshalJSON(b []byte) error {
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
