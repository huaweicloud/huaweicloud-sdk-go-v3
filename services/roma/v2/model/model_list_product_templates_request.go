package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListProductTemplatesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 每页显示条目数量，最大数量999，超过999后只返回999
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 产品模板ID
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 产品模板名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 产品模板状态 0-启用 1-停用
	Status *ListProductTemplatesRequestStatus `json:"status,omitempty" xml:"status"`

	// 创建用户名
	CreatedUserName *string `json:"created_user_name,omitempty" xml:"created_user_name"`

	// 创建时间起始，格式timestamp(ms)，使用UTC时区
	CreatedDateStart *int64 `json:"created_date_start,omitempty" xml:"created_date_start"`

	// 创建时间截止，格式timestamp(ms)，使用UTC时区
	CreatedDateEnd *int64 `json:"created_date_end,omitempty" xml:"created_date_end"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset *int32 `json:"offset,omitempty" xml:"offset"`
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
