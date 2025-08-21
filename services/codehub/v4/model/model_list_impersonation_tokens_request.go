package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListImpersonationTokensRequest Request Object
type ListImpersonationTokensRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`

	// **参数解释：** 状态 all 所有 active 活跃 inactive 非活跃。
	State *ListImpersonationTokensRequestState `json:"state,omitempty"`

	// **参数解释：** 检索内容
	Search *string `json:"search,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListImpersonationTokensRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImpersonationTokensRequest struct{}"
	}

	return strings.Join([]string{"ListImpersonationTokensRequest", string(data)}, " ")
}

type ListImpersonationTokensRequestState struct {
	value string
}

type ListImpersonationTokensRequestStateEnum struct {
	ALL      ListImpersonationTokensRequestState
	ACTIVE   ListImpersonationTokensRequestState
	INACTIVE ListImpersonationTokensRequestState
}

func GetListImpersonationTokensRequestStateEnum() ListImpersonationTokensRequestStateEnum {
	return ListImpersonationTokensRequestStateEnum{
		ALL: ListImpersonationTokensRequestState{
			value: "all",
		},
		ACTIVE: ListImpersonationTokensRequestState{
			value: "active",
		},
		INACTIVE: ListImpersonationTokensRequestState{
			value: "inactive",
		},
	}
}

func (c ListImpersonationTokensRequestState) Value() string {
	return c.value
}

func (c ListImpersonationTokensRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImpersonationTokensRequestState) UnmarshalJSON(b []byte) error {
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
