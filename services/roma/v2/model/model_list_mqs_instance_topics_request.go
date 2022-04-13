package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ListMqsInstanceTopicsRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
	// 应用名称。

	AppName *string `json:"app_name,omitempty"`
	// Topic名称。

	Name *string `json:"name,omitempty"`
	// 权限类型。 - all：发布+订阅 - pub：发布 - sub：订阅

	AccessPolicy *ListMqsInstanceTopicsRequestAccessPolicy `json:"access_policy,omitempty"`
	// 分页查询大小。默认查询所有的topic。

	Limit *string `json:"limit,omitempty"`
	// 分页查询的偏移量。默认值是0。

	Offset *string `json:"offset,omitempty"`
}

func (o ListMqsInstanceTopicsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMqsInstanceTopicsRequest struct{}"
	}

	return strings.Join([]string{"ListMqsInstanceTopicsRequest", string(data)}, " ")
}

type ListMqsInstanceTopicsRequestAccessPolicy struct {
	value string
}

type ListMqsInstanceTopicsRequestAccessPolicyEnum struct {
	ALL ListMqsInstanceTopicsRequestAccessPolicy
	PUB ListMqsInstanceTopicsRequestAccessPolicy
	SUB ListMqsInstanceTopicsRequestAccessPolicy
}

func GetListMqsInstanceTopicsRequestAccessPolicyEnum() ListMqsInstanceTopicsRequestAccessPolicyEnum {
	return ListMqsInstanceTopicsRequestAccessPolicyEnum{
		ALL: ListMqsInstanceTopicsRequestAccessPolicy{
			value: "all",
		},
		PUB: ListMqsInstanceTopicsRequestAccessPolicy{
			value: "pub",
		},
		SUB: ListMqsInstanceTopicsRequestAccessPolicy{
			value: "sub",
		},
	}
}

func (c ListMqsInstanceTopicsRequestAccessPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMqsInstanceTopicsRequestAccessPolicy) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
