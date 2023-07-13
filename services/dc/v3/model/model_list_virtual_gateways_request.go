package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListVirtualGatewaysRequest Request Object
type ListVirtualGatewaysRequest struct {

	// 每页返回的个数。 取值范围：1~2000。
	Limit *int32 `json:"limit,omitempty"`

	// 上一页最后一条资源记录的ID，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 显示字段列表
	Fields *[]string `json:"fields,omitempty"`

	// 返回结果按照升序(asc)或降序(desc)排列，默认为asc
	SortDir *[]ListVirtualGatewaysRequestSortDir `json:"sort_dir,omitempty"`

	// 排序字段。
	SortKey *string `json:"sort_key,omitempty"`

	// 根据资源ID过滤实例
	Id *[]string `json:"id,omitempty"`

	// 通过VPC-ID过虑虚拟网关实例
	VpcId *[]string `json:"vpc_id,omitempty"`
}

func (o ListVirtualGatewaysRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVirtualGatewaysRequest struct{}"
	}

	return strings.Join([]string{"ListVirtualGatewaysRequest", string(data)}, " ")
}

type ListVirtualGatewaysRequestSortDir struct {
	value string
}

type ListVirtualGatewaysRequestSortDirEnum struct {
	ASC  ListVirtualGatewaysRequestSortDir
	DESC ListVirtualGatewaysRequestSortDir
}

func GetListVirtualGatewaysRequestSortDirEnum() ListVirtualGatewaysRequestSortDirEnum {
	return ListVirtualGatewaysRequestSortDirEnum{
		ASC: ListVirtualGatewaysRequestSortDir{
			value: "asc",
		},
		DESC: ListVirtualGatewaysRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListVirtualGatewaysRequestSortDir) Value() string {
	return c.value
}

func (c ListVirtualGatewaysRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVirtualGatewaysRequestSortDir) UnmarshalJSON(b []byte) error {
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
