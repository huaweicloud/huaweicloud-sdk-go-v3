package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEndpointsRequest Request Object
type ListEndpointsRequest struct {

	// 偏移量，表示从此偏移量开始查询，偏移量不能小于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量，不能小于1或大于1000
	Limit *int32 `json:"limit,omitempty"`

	// 指定查询排序
	Sort *string `json:"sort,omitempty"`

	// 指定查询访问端点的类型
	Type *ListEndpointsRequestType `json:"type,omitempty"`

	// 指定查询访问端点的名称
	Name *string `json:"name,omitempty"`

	// 指定查询访问端点的vpcId
	VpcId *string `json:"vpc_id,omitempty"`

	// 指定查询访问端点的名称,模糊查询
	FuzzyName *string `json:"fuzzy_name,omitempty"`

	// 指定查询访问端点的SubnetId
	SubnetId *string `json:"subnet_id,omitempty"`
}

func (o ListEndpointsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointsRequest struct{}"
	}

	return strings.Join([]string{"ListEndpointsRequest", string(data)}, " ")
}

type ListEndpointsRequestType struct {
	value string
}

type ListEndpointsRequestTypeEnum struct {
	PRIVATE ListEndpointsRequestType
	PUBLIC  ListEndpointsRequestType
}

func GetListEndpointsRequestTypeEnum() ListEndpointsRequestTypeEnum {
	return ListEndpointsRequestTypeEnum{
		PRIVATE: ListEndpointsRequestType{
			value: "PRIVATE",
		},
		PUBLIC: ListEndpointsRequestType{
			value: "PUBLIC",
		},
	}
}

func (c ListEndpointsRequestType) Value() string {
	return c.value
}

func (c ListEndpointsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEndpointsRequestType) UnmarshalJSON(b []byte) error {
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
