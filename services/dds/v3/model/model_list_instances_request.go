package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListInstancesRequest struct {
	// 实例ID，可以调用“查询实例列表”接口获取。如果未申请实例，可以调用“创建实例”接口创建。

	Id *string `json:"id,omitempty"`
	// 实例名称。 如果name以“*”起始，表示按照“*”后面的值模糊匹配，否则，按照实际填写的name精确匹配查询。 - “*”为系统保留字符，不能只传入该字符。

	Name *string `json:"name,omitempty"`
	// 实例类型。 - 取值为“Sharding”，表示集群实例。 - 取值为“ReplicaSet”，表示副本集实例。 - 取值为“Single”，表示单节点实例。

	Mode *ListInstancesRequestMode `json:"mode,omitempty"`
	// 数据库版本类型。取值为“DDS-Community”。

	DatastoreType *ListInstancesRequestDatastoreType `json:"datastore_type,omitempty"`
	// 虚拟私有云ID。可登录虚拟私有云控制台界面，获取DDS实例所在虚拟私有云的ID。

	VpcId *string `json:"vpc_id,omitempty"`
	// 子网的网络ID。可登录虚拟私有云控制台界面，获取DDS实例所在虚拟私有云下子网的网络ID。

	SubnetId *string `json:"subnet_id,omitempty"`
	// 索引位置偏移量，表示从指定project ID下最新的实例创建时间开始，按时间的先后顺序偏移offset条数据后查询对应的实例信息。 取值大于或等于0。不传该参数时，查询偏移量默认为0，表示从最新的实例创建时间对应的实例开始查询。

	Offset *int32 `json:"offset,omitempty"`
	// 查询实例个数上限值。 取值范围：1~100。不传该参数时，默认查询前100条实例信息。

	Limit *int32 `json:"limit,omitempty"`
	// 根据实例标签键值对进行查询。{key}表示标签键，{value}表示标签值，最多包含20组。key不可以为空或重复，value可以为空。如果同时使用多个标签键值对进行查询，中间使用逗号分隔开，表示查询同时包含指定标签键值对的实例。

	Tags *string `json:"tags,omitempty"`
}

func (o ListInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesRequest", string(data)}, " ")
}

type ListInstancesRequestMode struct {
	value string
}

type ListInstancesRequestModeEnum struct {
	SHARDING    ListInstancesRequestMode
	REPLICA_SET ListInstancesRequestMode
	SINGLE      ListInstancesRequestMode
}

func GetListInstancesRequestModeEnum() ListInstancesRequestModeEnum {
	return ListInstancesRequestModeEnum{
		SHARDING: ListInstancesRequestMode{
			value: "Sharding",
		},
		REPLICA_SET: ListInstancesRequestMode{
			value: "ReplicaSet",
		},
		SINGLE: ListInstancesRequestMode{
			value: "Single",
		},
	}
}

func (c ListInstancesRequestMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesRequestMode) UnmarshalJSON(b []byte) error {
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

type ListInstancesRequestDatastoreType struct {
	value string
}

type ListInstancesRequestDatastoreTypeEnum struct {
	DDS_COMMUNITY ListInstancesRequestDatastoreType
	DDS_ENHANCED  ListInstancesRequestDatastoreType
}

func GetListInstancesRequestDatastoreTypeEnum() ListInstancesRequestDatastoreTypeEnum {
	return ListInstancesRequestDatastoreTypeEnum{
		DDS_COMMUNITY: ListInstancesRequestDatastoreType{
			value: "DDS-Community",
		},
		DDS_ENHANCED: ListInstancesRequestDatastoreType{
			value: "DDS-Enhanced",
		},
	}
}

func (c ListInstancesRequestDatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesRequestDatastoreType) UnmarshalJSON(b []byte) error {
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
