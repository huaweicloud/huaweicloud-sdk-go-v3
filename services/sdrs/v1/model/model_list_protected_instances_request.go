package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListProtectedInstancesRequest struct {

	// 保护组的ID，表示查询该保护组下的所有保护实例列表。
	ServerGroupId *string `json:"server_group_id,omitempty"`

	// 保护组的ID列表，格式为server_group_ids=['server_group_id1','server_group_id2',...,'server_group_idx']，请使用URL编码进行转换。返回“server_group_ids”中有效server_group_id的所有保护实例列表，无效的server_group_id会被忽略。支持查询最多30个server_group_id对应的所有保护实例列表。如果“server_group_id”和“server_group_ids”查询参数同时存在，“server_group_id”会被忽略。
	ServerGroupIds *string `json:"server_group_ids,omitempty"`

	// 保护实例的ID列表，格式为protected_instance_ids=['protected_instance_id1','protected_instance_id2',...,'protected_instance_idx']，请使用URL编码进行转换。返回“protected_instance_ids”中有效protected_instance_id的所有保护实例列表，无效的protected_instance_id会被忽略。支持查询最多30个protected_instance_id对应的所有保护实例列表。如果“server_group_id”或者“server_group_ids”查询参数存在时，“protected_instance_ids”会被忽略。
	ProtectedInstanceIds *string `json:"protected_instance_ids,omitempty"`

	// 每次请求返回结果个数限制，取值范围为[0,1000]的正整数，默认值为1000。
	Limit *int32 `json:"limit,omitempty"`

	// 每次请求开始的下标，即偏移量，默认值为0。offset必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 保护实例状态。
	Status *string `json:"status,omitempty"`

	// 保护实例的名称。支持模糊查询。
	Name *string `json:"name,omitempty"`

	// 查询场景类型。status_abnormal：表示查询异常状态的保护实例列表。general或空时：该参数不生效。
	QueryType *ListProtectedInstancesRequestQueryType `json:"query_type,omitempty"`

	// 保护实例所在的保护组的当前生产站点可用区。
	AvailabilityZone *string `json:"availability_zone,omitempty"`
}

func (o ListProtectedInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectedInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListProtectedInstancesRequest", string(data)}, " ")
}

type ListProtectedInstancesRequestQueryType struct {
	value string
}

type ListProtectedInstancesRequestQueryTypeEnum struct {
	STATUS_ABNORMAL ListProtectedInstancesRequestQueryType
	GENERAL         ListProtectedInstancesRequestQueryType
}

func GetListProtectedInstancesRequestQueryTypeEnum() ListProtectedInstancesRequestQueryTypeEnum {
	return ListProtectedInstancesRequestQueryTypeEnum{
		STATUS_ABNORMAL: ListProtectedInstancesRequestQueryType{
			value: "status_abnormal",
		},
		GENERAL: ListProtectedInstancesRequestQueryType{
			value: "general",
		},
	}
}

func (c ListProtectedInstancesRequestQueryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProtectedInstancesRequestQueryType) UnmarshalJSON(b []byte) error {
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
