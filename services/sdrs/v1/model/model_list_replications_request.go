package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListReplicationsRequest struct {

	// 保护组的ID。
	ServerGroupId *string `json:"server_group_id,omitempty" xml:"server_group_id"`

	// 保护组的ID列表，格式为server_group_ids=['server_group_id1','server_group_id2',...,'server_group_idx']，请使用URL编码进行转换。返回“server_group_ids”中有效server_group_id的复制对列表，无效的server_group_id会被忽略。支持查询最多30个server_group_id对应的复制对列表。如果“server_group_id”和“server_group_ids”查询参数同时存在，“server_group_id”会被忽略。
	ServerGroupIds *string `json:"server_group_ids,omitempty" xml:"server_group_ids"`

	// 保护实例的ID。
	ProtectedInstanceId *string `json:"protected_instance_id,omitempty" xml:"protected_instance_id"`

	// 保护实例的ID列表，格式为protected_instance_ids=['protected_instance_id1','protected_instance_id2',...,'protected_instance_idx']，请使用URL编码进行转换。返回“protected_instance_ids”中有效protected_instance_id的复制对列表，无效的protected_instance_id会被忽略。支持查询最多30个protected_instance_id对应的复制对列表。如果“protected_instance_id”和“protected_instance_ids”查询参数同时存在，“protected_instance_id”会被忽略。
	ProtectedInstanceIds *string `json:"protected_instance_ids,omitempty" xml:"protected_instance_ids"`

	// 复制对的名称。支持模糊查询。
	Name *string `json:"name,omitempty" xml:"name"`

	// 复制对的状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 每次请求返回结果个数限制，取值范围为[0,1000]的正整数，默认值为1000。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 每次请求开始的下标，即偏移量，默认值为0。offset必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 查询场景类型。如需查询异常状态的复制对列表，query_type的值为“status_abnormal”。否则，query_type取值为空或“general”。
	QueryType *ListReplicationsRequestQueryType `json:"query_type,omitempty" xml:"query_type"`

	// 复制对所在的保护组的当前生产站点可用区。
	AvailabilityZone *string `json:"availability_zone,omitempty" xml:"availability_zone"`
}

func (o ListReplicationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReplicationsRequest struct{}"
	}

	return strings.Join([]string{"ListReplicationsRequest", string(data)}, " ")
}

type ListReplicationsRequestQueryType struct {
	value string
}

type ListReplicationsRequestQueryTypeEnum struct {
	STATUS_ABNORMAL ListReplicationsRequestQueryType
	GENERAL         ListReplicationsRequestQueryType
}

func GetListReplicationsRequestQueryTypeEnum() ListReplicationsRequestQueryTypeEnum {
	return ListReplicationsRequestQueryTypeEnum{
		STATUS_ABNORMAL: ListReplicationsRequestQueryType{
			value: "status_abnormal",
		},
		GENERAL: ListReplicationsRequestQueryType{
			value: "general",
		},
	}
}

func (c ListReplicationsRequestQueryType) Value() string {
	return c.value
}

func (c ListReplicationsRequestQueryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListReplicationsRequestQueryType) UnmarshalJSON(b []byte) error {
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
