package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

//
type VaultResourceInstancesReq struct {
	// 不包含任意一个标签，该字段为true时查询所有不带标签的资源，此时忽略 “tags”、“tags_any”、“not_tags”、“not_tags_any”字段。

	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`
	// 包含标签。  tags不允许为空列表。  tags中最多包含10个key。  tags中key不允许重复。  tags中多个key之间是“与”的关系。  结果返回包含所有标签的资源列表，key之间是与的关系，key-value结构中value是或的关系。  无过滤条件时返回全量数据。

	Tags *[]TagsReq `json:"tags,omitempty"`
	// 包含任一标签。  tags不允许为空列表。  tags中最多包含10个key。  tags中key不允许重复。  结果返回包含任一标签的资源列表，key之间是或的关系，key-value结构中value是或的关系。  无过滤条件时返回全量数据。

	TagsAny *[]TagsReq `json:"tags_any,omitempty"`
	// 不包含标签。  tags不允许为空列表。  tags中最多包含10个key。  tags中key不允许重复。  结果返回不包含所有标签的资源列表，key之间是与的关系，key-value结构中value是或的关系。  无过滤条件时返回全量数据。

	NotTags *[]TagsReq `json:"not_tags,omitempty"`
	// 不包含任一标签。  tags不允许为空列表。  tags中最多包含10个key。  tags中key不允许重复。  结果返回不包含任一标签的资源列表，key之间是或的关系，key-value结构张value是或的关系。  无过滤条件时返回全量数据。

	NotTagsAny *[]TagsReq `json:"not_tags_any,omitempty"`
	// 仅op_service权限可以使用此字段做资源实例过滤条件。  目前TMS调用时只包含一个tag结构体。  * key： _sys_enterprise_project_id  * values：企业项目id列表  目前TMS调用时，key下面只包含一个value，0表示默认企业项目。  sys_tags和租户标签过滤条件（tags、tags_any、not_tags、not_tags_any）不能同时使用。  无sys_tags时按照tag接口处理，无tag过滤条件时返回全量数据。  sys_tags不能为空列表

	SysTags *[]SysTags `json:"sys_tags,omitempty"`
	// 查询记录数（action为count时无此参数）如果action为filter时，默认为1000，limit最小值为1，limit最大值为1000, 不在范围内报错。返回的结果中记录数不超过limit。

	Limit *string `json:"limit,omitempty"`
	// 索引位置（action为count时无此参数）如果action为filter时，默认为0，offset最小值为0。返回的结果中第一条记录为符合查询条件的第offset+1条记录。

	Offset *string `json:"offset,omitempty"`
	// 操作标识取值范围为：\"filter\", \"count\"。如果是filter就是分页查询，如果是count只需按照条件将总条数返回即可

	Action string `json:"action"`
	// 资源本身支持的查询条件。  matches不允许为空列表。  matches中key不允许重复。

	Matches *[]Match `json:"matches,omitempty"`
	// 云类型

	CloudType *VaultResourceInstancesReqCloudType `json:"cloud_type,omitempty"`
	// 资源类型

	ObjectType *VaultResourceInstancesReqObjectType `json:"object_type,omitempty"`
}

func (o VaultResourceInstancesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultResourceInstancesReq struct{}"
	}

	return strings.Join([]string{"VaultResourceInstancesReq", string(data)}, " ")
}

type VaultResourceInstancesReqCloudType struct {
	value string
}

type VaultResourceInstancesReqCloudTypeEnum struct {
	PUBLIC VaultResourceInstancesReqCloudType
	HYBRID VaultResourceInstancesReqCloudType
}

func GetVaultResourceInstancesReqCloudTypeEnum() VaultResourceInstancesReqCloudTypeEnum {
	return VaultResourceInstancesReqCloudTypeEnum{
		PUBLIC: VaultResourceInstancesReqCloudType{
			value: "public",
		},
		HYBRID: VaultResourceInstancesReqCloudType{
			value: " hybrid",
		},
	}
}

func (c VaultResourceInstancesReqCloudType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VaultResourceInstancesReqCloudType) UnmarshalJSON(b []byte) error {
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

type VaultResourceInstancesReqObjectType struct {
	value string
}

type VaultResourceInstancesReqObjectTypeEnum struct {
	SERVER VaultResourceInstancesReqObjectType
	DISK   VaultResourceInstancesReqObjectType
}

func GetVaultResourceInstancesReqObjectTypeEnum() VaultResourceInstancesReqObjectTypeEnum {
	return VaultResourceInstancesReqObjectTypeEnum{
		SERVER: VaultResourceInstancesReqObjectType{
			value: "server",
		},
		DISK: VaultResourceInstancesReqObjectType{
			value: "disk",
		},
	}
}

func (c VaultResourceInstancesReqObjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VaultResourceInstancesReqObjectType) UnmarshalJSON(b []byte) error {
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
