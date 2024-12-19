package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPeerLinksRequest Request Object
type ListPeerLinksRequest struct {

	// 全域接入网关ID
	GlobalDcGatewayId string `json:"global_dc_gateway_id"`

	// 每页返回的个数。 取值范围：1~2000。
	Limit *int32 `json:"limit,omitempty"`

	// 分页偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 上一页最后一条资源记录的ID，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 分页参数
	PageReverse *bool `json:"page_reverse,omitempty"`

	// 显示字段列表
	Fields *[]string `json:"fields,omitempty"`

	// show response ext-fields
	ExtFields *[]string `json:"ext_fields,omitempty"`

	// 排序字段。
	SortKey *string `json:"sort_key,omitempty"`

	// 返回结果按照升序(asc)或降序(desc)排列，默认为asc
	SortDir *[]ListPeerLinksRequestSortDir `json:"sort_dir,omitempty"`

	// 根据资源ID过滤实例
	Id *[]string `json:"id,omitempty"`

	// 根据名字过滤查询，可查询多个名字。
	Name *[]string `json:"name,omitempty"`
}

func (o ListPeerLinksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPeerLinksRequest struct{}"
	}

	return strings.Join([]string{"ListPeerLinksRequest", string(data)}, " ")
}

type ListPeerLinksRequestSortDir struct {
	value string
}

type ListPeerLinksRequestSortDirEnum struct {
	ASC  ListPeerLinksRequestSortDir
	DESC ListPeerLinksRequestSortDir
}

func GetListPeerLinksRequestSortDirEnum() ListPeerLinksRequestSortDirEnum {
	return ListPeerLinksRequestSortDirEnum{
		ASC: ListPeerLinksRequestSortDir{
			value: "asc",
		},
		DESC: ListPeerLinksRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListPeerLinksRequestSortDir) Value() string {
	return c.value
}

func (c ListPeerLinksRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPeerLinksRequestSortDir) UnmarshalJSON(b []byte) error {
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
