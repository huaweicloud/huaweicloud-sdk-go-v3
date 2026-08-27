package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListModelGroupsRequest Request Object
type ListModelGroupsRequest struct {

	// 偏移量，默认0。
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小，默认20。
	Limit *int32 `json:"limit,omitempty"`

	// 分组名称（模糊匹配）。
	Name *string `json:"name,omitempty"`

	// 分组状态（draft-草稿，active-活跃，inactive-停用，deprecated-废弃）。
	Status *ListModelGroupsRequestStatus `json:"status,omitempty"`
}

func (o ListModelGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListModelGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListModelGroupsRequest", string(data)}, " ")
}

type ListModelGroupsRequestStatus struct {
	value string
}

type ListModelGroupsRequestStatusEnum struct {
	DRAFT      ListModelGroupsRequestStatus
	ACTIVE     ListModelGroupsRequestStatus
	INACTIVE   ListModelGroupsRequestStatus
	DEPRECATED ListModelGroupsRequestStatus
}

func GetListModelGroupsRequestStatusEnum() ListModelGroupsRequestStatusEnum {
	return ListModelGroupsRequestStatusEnum{
		DRAFT: ListModelGroupsRequestStatus{
			value: "draft",
		},
		ACTIVE: ListModelGroupsRequestStatus{
			value: "active",
		},
		INACTIVE: ListModelGroupsRequestStatus{
			value: "inactive",
		},
		DEPRECATED: ListModelGroupsRequestStatus{
			value: "deprecated",
		},
	}
}

func (c ListModelGroupsRequestStatus) Value() string {
	return c.value
}

func (c ListModelGroupsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListModelGroupsRequestStatus) UnmarshalJSON(b []byte) error {
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
