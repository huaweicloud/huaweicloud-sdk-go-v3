package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListChildProxyRepositoriesListResponse Response Object
type ListChildProxyRepositoriesListResponse struct {

	// **参数解释**： 请求成功或失败状态。 **取值范围**： - success：请求成功。 - error：请求失败。
	Status *ListChildProxyRepositoriesListResponseStatus `json:"status,omitempty"`

	// **参数解释**： 请求ID，当前请求唯一标识。 **取值范围**： 数字及中划线（-）组成的字符串。
	TraceId *string `json:"trace_id,omitempty"`

	// **参数解释**: 仓库列表。 **取值范围**: 不涉及。
	Result         *[]RepositoryDov5 `json:"result,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListChildProxyRepositoriesListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListChildProxyRepositoriesListResponse struct{}"
	}

	return strings.Join([]string{"ListChildProxyRepositoriesListResponse", string(data)}, " ")
}

type ListChildProxyRepositoriesListResponseStatus struct {
	value string
}

type ListChildProxyRepositoriesListResponseStatusEnum struct {
	SUCCESS ListChildProxyRepositoriesListResponseStatus
	ERROR   ListChildProxyRepositoriesListResponseStatus
}

func GetListChildProxyRepositoriesListResponseStatusEnum() ListChildProxyRepositoriesListResponseStatusEnum {
	return ListChildProxyRepositoriesListResponseStatusEnum{
		SUCCESS: ListChildProxyRepositoriesListResponseStatus{
			value: "success",
		},
		ERROR: ListChildProxyRepositoriesListResponseStatus{
			value: "error",
		},
	}
}

func (c ListChildProxyRepositoriesListResponseStatus) Value() string {
	return c.value
}

func (c ListChildProxyRepositoriesListResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListChildProxyRepositoriesListResponseStatus) UnmarshalJSON(b []byte) error {
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
