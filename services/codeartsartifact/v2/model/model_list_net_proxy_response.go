package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListNetProxyResponse Response Object
type ListNetProxyResponse struct {

	// **参数解释**： 请求成功或失败状态。 **取值范围**： - success：请求成功。 - error：请求失败。
	Status *ListNetProxyResponseStatus `json:"status,omitempty"`

	// **参数解释**： 请求ID，当前请求唯一标识。 **取值范围**： 数字及中划线（-）组成的字符串。
	TraceId *string `json:"trace_id,omitempty"`

	// **参数解释**: 网络代理列表。 **取值范围**: 不涉及。
	Result         *[]NetProxyModelV5 `json:"result,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListNetProxyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNetProxyResponse struct{}"
	}

	return strings.Join([]string{"ListNetProxyResponse", string(data)}, " ")
}

type ListNetProxyResponseStatus struct {
	value string
}

type ListNetProxyResponseStatusEnum struct {
	SUCCESS ListNetProxyResponseStatus
	ERROR   ListNetProxyResponseStatus
}

func GetListNetProxyResponseStatusEnum() ListNetProxyResponseStatusEnum {
	return ListNetProxyResponseStatusEnum{
		SUCCESS: ListNetProxyResponseStatus{
			value: "success",
		},
		ERROR: ListNetProxyResponseStatus{
			value: "error",
		},
	}
}

func (c ListNetProxyResponseStatus) Value() string {
	return c.value
}

func (c ListNetProxyResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNetProxyResponseStatus) UnmarshalJSON(b []byte) error {
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
