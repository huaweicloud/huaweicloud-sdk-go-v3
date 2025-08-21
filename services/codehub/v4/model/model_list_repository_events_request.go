package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ListRepositoryEventsRequest Request Object
type ListRepositoryEventsRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 操作者id。
	AuthorId *int32 `json:"author_id,omitempty"`

	// **参数解释：** 动态类型。 **约束限制：** - all，表示全部。 - push，表示推送。
	Filter *ListRepositoryEventsRequestFilter `json:"filter,omitempty"`

	// **参数解释：** 开始日期。
	Before *sdktime.SdkTime `json:"before,omitempty"`

	// **参数解释：** 结束日期。
	After *sdktime.SdkTime `json:"after,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListRepositoryEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryEventsRequest struct{}"
	}

	return strings.Join([]string{"ListRepositoryEventsRequest", string(data)}, " ")
}

type ListRepositoryEventsRequestFilter struct {
	value string
}

type ListRepositoryEventsRequestFilterEnum struct {
	ALL  ListRepositoryEventsRequestFilter
	PUSH ListRepositoryEventsRequestFilter
}

func GetListRepositoryEventsRequestFilterEnum() ListRepositoryEventsRequestFilterEnum {
	return ListRepositoryEventsRequestFilterEnum{
		ALL: ListRepositoryEventsRequestFilter{
			value: "all",
		},
		PUSH: ListRepositoryEventsRequestFilter{
			value: "push",
		},
	}
}

func (c ListRepositoryEventsRequestFilter) Value() string {
	return c.value
}

func (c ListRepositoryEventsRequestFilter) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRepositoryEventsRequestFilter) UnmarshalJSON(b []byte) error {
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
