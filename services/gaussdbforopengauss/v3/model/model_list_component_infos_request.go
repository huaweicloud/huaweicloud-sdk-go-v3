package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListComponentInfosRequest Request Object
type ListComponentInfosRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100
	Limit *int32 `json:"limit,omitempty"`

	// 组件类型，过滤拿到需要的组件类型的组件信息，默认为ALL，传参数会查询对应组件信息。 枚举值：   \"ALL\": 查询所有组件类型。   \"CN\": 查询CN组件类型。   \"DN\": 查询DN组件类型。   \"CM\": 查询CMS组件类型。   \"GTM\": 查询GTM组件类型。   \"ETCD\": 查询ETCD组件类型。
	ComponentType *ListComponentInfosRequestComponentType `json:"component_type,omitempty"`

	// 主组件所在可用区编号，筛选符合条件的组件，默认为ALL，查询实例所有可用区上的节点的组件信息。 当调用接口传入可用区编号时：   相对于DN组件，会查询出的DN分片中的主组件在该可用区上的这个分片的所有副本的组件信息。   相对于CN组件，CN组件没有主备关系，会查询出该可用区上的CN组件信息。   相对于其他组件，会查询该可用区上有没有某个组件类型的主组件，有则会返回该组件类型的所有组件信息，没有则不返回该组件类型的信息。
	AvailabilityZoneId *string `json:"availability_zone_id,omitempty"`
}

func (o ListComponentInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComponentInfosRequest struct{}"
	}

	return strings.Join([]string{"ListComponentInfosRequest", string(data)}, " ")
}

type ListComponentInfosRequestComponentType struct {
	value string
}

type ListComponentInfosRequestComponentTypeEnum struct {
	ALL  ListComponentInfosRequestComponentType
	CN   ListComponentInfosRequestComponentType
	DN   ListComponentInfosRequestComponentType
	CM   ListComponentInfosRequestComponentType
	GTM  ListComponentInfosRequestComponentType
	ETCD ListComponentInfosRequestComponentType
}

func GetListComponentInfosRequestComponentTypeEnum() ListComponentInfosRequestComponentTypeEnum {
	return ListComponentInfosRequestComponentTypeEnum{
		ALL: ListComponentInfosRequestComponentType{
			value: "ALL",
		},
		CN: ListComponentInfosRequestComponentType{
			value: "CN",
		},
		DN: ListComponentInfosRequestComponentType{
			value: "DN",
		},
		CM: ListComponentInfosRequestComponentType{
			value: "CM",
		},
		GTM: ListComponentInfosRequestComponentType{
			value: "GTM",
		},
		ETCD: ListComponentInfosRequestComponentType{
			value: "ETCD",
		},
	}
}

func (c ListComponentInfosRequestComponentType) Value() string {
	return c.value
}

func (c ListComponentInfosRequestComponentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListComponentInfosRequestComponentType) UnmarshalJSON(b []byte) error {
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
