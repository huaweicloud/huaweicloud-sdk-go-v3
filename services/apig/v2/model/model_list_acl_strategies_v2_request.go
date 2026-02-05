package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAclStrategiesV2Request Request Object
type ListAclStrategiesV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量，条目数量小于等于0时，自动转换为20，条目数量大于500时，自动转换为500
	Limit *int32 `json:"limit,omitempty"`

	// ACL策略编号。
	Id *string `json:"id,omitempty"`

	// ACL策略名称。
	Name *string `json:"name,omitempty"`

	// 类型。 - PERMIT：白名单类型 - DENY：黑名单类型
	AclType *ListAclStrategiesV2RequestAclType `json:"acl_type,omitempty"`

	// 作用的对象类型： - IP - DOMAIN
	EntityType *string `json:"entity_type,omitempty"`

	// 指定需要精确匹配查找的参数名称，目前仅支持name
	PreciseSearch *string `json:"precise_search,omitempty"`
}

func (o ListAclStrategiesV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAclStrategiesV2Request struct{}"
	}

	return strings.Join([]string{"ListAclStrategiesV2Request", string(data)}, " ")
}

type ListAclStrategiesV2RequestAclType struct {
	value string
}

type ListAclStrategiesV2RequestAclTypeEnum struct {
	PERMIT ListAclStrategiesV2RequestAclType
	DENY   ListAclStrategiesV2RequestAclType
}

func GetListAclStrategiesV2RequestAclTypeEnum() ListAclStrategiesV2RequestAclTypeEnum {
	return ListAclStrategiesV2RequestAclTypeEnum{
		PERMIT: ListAclStrategiesV2RequestAclType{
			value: "PERMIT",
		},
		DENY: ListAclStrategiesV2RequestAclType{
			value: "DENY",
		},
	}
}

func (c ListAclStrategiesV2RequestAclType) Value() string {
	return c.value
}

func (c ListAclStrategiesV2RequestAclType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAclStrategiesV2RequestAclType) UnmarshalJSON(b []byte) error {
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
