package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListServersByHostIdRequest Request Object
type ListServersByHostIdRequest struct {

	// 云办公主机id。
	HostId string `json:"host_id"`

	// 排序字段名称，需要结合sort_type字段一起使用。 - vcpu CPU核数 - memory 内存大小
	SortField *ListServersByHostIdRequestSortField `json:"sort_field,omitempty"`

	// 排序类型，默认升序，需要结合sort_field字段一起使用。 - ASC 升序。 - DESC 降序。
	SortType *ListServersByHostIdRequestSortType `json:"sort_type,omitempty"`

	// 每页显示的数量。
	Limit *int32 `json:"limit,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListServersByHostIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServersByHostIdRequest struct{}"
	}

	return strings.Join([]string{"ListServersByHostIdRequest", string(data)}, " ")
}

type ListServersByHostIdRequestSortField struct {
	value string
}

type ListServersByHostIdRequestSortFieldEnum struct {
	VCPU   ListServersByHostIdRequestSortField
	MEMORY ListServersByHostIdRequestSortField
}

func GetListServersByHostIdRequestSortFieldEnum() ListServersByHostIdRequestSortFieldEnum {
	return ListServersByHostIdRequestSortFieldEnum{
		VCPU: ListServersByHostIdRequestSortField{
			value: "vcpu",
		},
		MEMORY: ListServersByHostIdRequestSortField{
			value: "memory",
		},
	}
}

func (c ListServersByHostIdRequestSortField) Value() string {
	return c.value
}

func (c ListServersByHostIdRequestSortField) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServersByHostIdRequestSortField) UnmarshalJSON(b []byte) error {
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

type ListServersByHostIdRequestSortType struct {
	value string
}

type ListServersByHostIdRequestSortTypeEnum struct {
	ASC  ListServersByHostIdRequestSortType
	DESC ListServersByHostIdRequestSortType
}

func GetListServersByHostIdRequestSortTypeEnum() ListServersByHostIdRequestSortTypeEnum {
	return ListServersByHostIdRequestSortTypeEnum{
		ASC: ListServersByHostIdRequestSortType{
			value: "ASC",
		},
		DESC: ListServersByHostIdRequestSortType{
			value: "DESC",
		},
	}
}

func (c ListServersByHostIdRequestSortType) Value() string {
	return c.value
}

func (c ListServersByHostIdRequestSortType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServersByHostIdRequestSortType) UnmarshalJSON(b []byte) error {
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
