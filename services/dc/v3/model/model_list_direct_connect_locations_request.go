package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDirectConnectLocationsRequest Request Object
type ListDirectConnectLocationsRequest struct {

	// 每页返回的个数。 取值范围：1~2000。
	Limit *int32 `json:"limit,omitempty"`

	// 上一页最后一条资源记录的ID，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 排序字段。
	SortKey *string `json:"sort_key,omitempty"`

	// 返回结果按照升序(asc)或降序(desc)排列，默认为asc
	SortDir *[]ListDirectConnectLocationsRequestSortDir `json:"sort_dir,omitempty"`

	// 根据资源ID过滤实例
	Id *[]string `json:"id,omitempty"`

	// 根据名字过滤查询，可查询多个名字。
	Name *[]string `json:"name,omitempty"`

	// 根椐资源状态过滤实例
	Status *[]string `json:"status,omitempty"`
}

func (o ListDirectConnectLocationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDirectConnectLocationsRequest struct{}"
	}

	return strings.Join([]string{"ListDirectConnectLocationsRequest", string(data)}, " ")
}

type ListDirectConnectLocationsRequestSortDir struct {
	value string
}

type ListDirectConnectLocationsRequestSortDirEnum struct {
	ASC  ListDirectConnectLocationsRequestSortDir
	DESC ListDirectConnectLocationsRequestSortDir
}

func GetListDirectConnectLocationsRequestSortDirEnum() ListDirectConnectLocationsRequestSortDirEnum {
	return ListDirectConnectLocationsRequestSortDirEnum{
		ASC: ListDirectConnectLocationsRequestSortDir{
			value: "asc",
		},
		DESC: ListDirectConnectLocationsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListDirectConnectLocationsRequestSortDir) Value() string {
	return c.value
}

func (c ListDirectConnectLocationsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDirectConnectLocationsRequestSortDir) UnmarshalJSON(b []byte) error {
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