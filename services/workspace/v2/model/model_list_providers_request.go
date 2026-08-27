package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListProvidersRequest Request Object
type ListProvidersRequest struct {

	// 每页数量，默认10，最大100。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，默认0。
	Offset *int32 `json:"offset,omitempty"`

	// 按分组筛选（关联查询）。
	GroupId *string `json:"group_id,omitempty"`

	// 状态筛选（connected-已连接，disconnected-已断开，unverified-未验证）。
	Status *ListProvidersRequestStatus `json:"status,omitempty"`

	// 供应商类型筛选。
	ProviderType *string `json:"provider_type,omitempty"`

	// 名称模糊搜索。
	Name *string `json:"name,omitempty"`
}

func (o ListProvidersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProvidersRequest struct{}"
	}

	return strings.Join([]string{"ListProvidersRequest", string(data)}, " ")
}

type ListProvidersRequestStatus struct {
	value string
}

type ListProvidersRequestStatusEnum struct {
	CONNECTED    ListProvidersRequestStatus
	DISCONNECTED ListProvidersRequestStatus
	UNVERIFIED   ListProvidersRequestStatus
}

func GetListProvidersRequestStatusEnum() ListProvidersRequestStatusEnum {
	return ListProvidersRequestStatusEnum{
		CONNECTED: ListProvidersRequestStatus{
			value: "connected",
		},
		DISCONNECTED: ListProvidersRequestStatus{
			value: "disconnected",
		},
		UNVERIFIED: ListProvidersRequestStatus{
			value: "unverified",
		},
	}
}

func (c ListProvidersRequestStatus) Value() string {
	return c.value
}

func (c ListProvidersRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProvidersRequestStatus) UnmarshalJSON(b []byte) error {
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
