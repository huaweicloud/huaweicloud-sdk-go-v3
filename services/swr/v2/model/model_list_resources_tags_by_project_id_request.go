package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourcesTagsByProjectIdRequest Request Object
type ListResourcesTagsByProjectIdRequest struct {

	// 资源类型 取值范围： - instances：Instance实例 - namespaces：命名空间 - repositories：仓库
	ResourceType ListResourcesTagsByProjectIdRequestResourceType `json:"resource_type"`

	// 起始索引，默认值为0。**注意：offset和limit参数需要配套使用。**
	Offset *int32 `json:"offset,omitempty"`

	// 返回条数，默认为200，最小值为1，最大值为200。**注意：offset和limit参数需要配套使用。**
	Limit *int32 `json:"limit,omitempty"`

	// 当resource_type为repositories或namespaces时，支持通过instance_id进行过滤查询，用于获取指定注册表实例下对应资源类型的标签列表。
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o ListResourcesTagsByProjectIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesTagsByProjectIdRequest struct{}"
	}

	return strings.Join([]string{"ListResourcesTagsByProjectIdRequest", string(data)}, " ")
}

type ListResourcesTagsByProjectIdRequestResourceType struct {
	value string
}

type ListResourcesTagsByProjectIdRequestResourceTypeEnum struct {
	INSTANCES    ListResourcesTagsByProjectIdRequestResourceType
	NAMESPACES   ListResourcesTagsByProjectIdRequestResourceType
	REPOSITORIES ListResourcesTagsByProjectIdRequestResourceType
}

func GetListResourcesTagsByProjectIdRequestResourceTypeEnum() ListResourcesTagsByProjectIdRequestResourceTypeEnum {
	return ListResourcesTagsByProjectIdRequestResourceTypeEnum{
		INSTANCES: ListResourcesTagsByProjectIdRequestResourceType{
			value: "instances",
		},
		NAMESPACES: ListResourcesTagsByProjectIdRequestResourceType{
			value: "namespaces",
		},
		REPOSITORIES: ListResourcesTagsByProjectIdRequestResourceType{
			value: "repositories",
		},
	}
}

func (c ListResourcesTagsByProjectIdRequestResourceType) Value() string {
	return c.value
}

func (c ListResourcesTagsByProjectIdRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourcesTagsByProjectIdRequestResourceType) UnmarshalJSON(b []byte) error {
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
