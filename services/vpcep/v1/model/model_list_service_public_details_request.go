package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListServicePublicDetailsRequest Request Object
type ListServicePublicDetailsRequest struct {

	// 发送的实体的MIME类型。推荐用户默认使用application/json， 如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。
	ContentType *string `json:"Content-Type,omitempty"`

	// 查询返回公共的终端节点服务数量限制，即每页返回的个数。 取值范围：0~1000，取值一般为10，20或者50，默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量。 偏移量为一个大于0小于终端节点服务总个数的整数， 表示从偏移量后面的终端节点服务开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 公共终端节点服务的名称，支持大小写以及模糊匹配。
	EndpointServiceName *string `json:"endpoint_service_name,omitempty"`

	// 公共终端节点服务的ID，唯一标识。
	Id *string `json:"id,omitempty"`

	// 查询结果中终端节点服务列表的排序字段，取值为：  - create_at：终端节点服务的创建时间  - update_at：终端节点服务的更新时间 默认值为create_at。
	SortKey *ListServicePublicDetailsRequestSortKey `json:"sort_key,omitempty"`

	// 查询结果中终端节点服务列表的排序方式，取值为：  - desc：降序排序  - asc：升序排序 默认值为desc。
	SortDir *ListServicePublicDetailsRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListServicePublicDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServicePublicDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListServicePublicDetailsRequest", string(data)}, " ")
}

type ListServicePublicDetailsRequestSortKey struct {
	value string
}

type ListServicePublicDetailsRequestSortKeyEnum struct {
	CREATE_AT ListServicePublicDetailsRequestSortKey
	UPDATE_AT ListServicePublicDetailsRequestSortKey
}

func GetListServicePublicDetailsRequestSortKeyEnum() ListServicePublicDetailsRequestSortKeyEnum {
	return ListServicePublicDetailsRequestSortKeyEnum{
		CREATE_AT: ListServicePublicDetailsRequestSortKey{
			value: "create_at",
		},
		UPDATE_AT: ListServicePublicDetailsRequestSortKey{
			value: "update_at",
		},
	}
}

func (c ListServicePublicDetailsRequestSortKey) Value() string {
	return c.value
}

func (c ListServicePublicDetailsRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServicePublicDetailsRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListServicePublicDetailsRequestSortDir struct {
	value string
}

type ListServicePublicDetailsRequestSortDirEnum struct {
	ASC  ListServicePublicDetailsRequestSortDir
	DESC ListServicePublicDetailsRequestSortDir
}

func GetListServicePublicDetailsRequestSortDirEnum() ListServicePublicDetailsRequestSortDirEnum {
	return ListServicePublicDetailsRequestSortDirEnum{
		ASC: ListServicePublicDetailsRequestSortDir{
			value: "asc",
		},
		DESC: ListServicePublicDetailsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListServicePublicDetailsRequestSortDir) Value() string {
	return c.value
}

func (c ListServicePublicDetailsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServicePublicDetailsRequestSortDir) UnmarshalJSON(b []byte) error {
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
