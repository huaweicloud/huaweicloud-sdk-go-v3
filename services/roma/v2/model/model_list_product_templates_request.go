package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListProductTemplatesRequest Request Object
type ListProductTemplatesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 每页显示条目数量，最大数量999，超过999后只返回999
	Limit *int32 `json:"limit,omitempty"`

	// 产品模板ID
	Id *int32 `json:"id,omitempty"`

	// 产品模板名称
	Name *string `json:"name,omitempty"`

	// 产品模板状态 0-启用 1-停用
	Status *ListProductTemplatesRequestStatus `json:"status,omitempty"`

	// 创建用户名
	CreatedUserName *string `json:"created_user_name,omitempty"`

	// 创建时间起始，格式timestamp(ms)，使用UTC时区
	CreatedDateStart *int64 `json:"created_date_start,omitempty"`

	// 创建时间截止，格式timestamp(ms)，使用UTC时区
	CreatedDateEnd *int64 `json:"created_date_end,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListProductTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListProductTemplatesRequest", string(data)}, " ")
}

type ListProductTemplatesRequestStatus struct {
	value int32
}

type ListProductTemplatesRequestStatusEnum struct {
	E_0 ListProductTemplatesRequestStatus
	E_1 ListProductTemplatesRequestStatus
}

func GetListProductTemplatesRequestStatusEnum() ListProductTemplatesRequestStatusEnum {
	return ListProductTemplatesRequestStatusEnum{
		E_0: ListProductTemplatesRequestStatus{
			value: 0,
		}, E_1: ListProductTemplatesRequestStatus{
			value: 1,
		},
	}
}

func (c ListProductTemplatesRequestStatus) Value() int32 {
	return c.value
}

func (c ListProductTemplatesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProductTemplatesRequestStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
