package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListGlobalEipsRequest Request Object
type ListGlobalEipsRequest struct {

	// 每页返回的个数。 取值范围：1~2000。
	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	// 上一页最后一条资源记录的ID，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	PageReverse *bool `json:"page_reverse,omitempty"`

	// 显示字段列表
	Fields *[]string `json:"fields,omitempty"`

	// show response ext-fields
	ExtFields *[]string `json:"ext_fields,omitempty"`

	// 排序字段。
	SortKey *string `json:"sort_key,omitempty"`

	// 返回结果按照升序(asc)或降序(desc)排列，默认为asc
	SortDir *[]ListGlobalEipsRequestSortDir `json:"sort_dir,omitempty"`

	// 互联网关ID
	ConnectGatewayId string `json:"connect_gateway_id"`

	// 根椐资源状态过淲实例
	Status *[]string `json:"status,omitempty"`

	// 全局弹性IP的ID
	GlobalEipId *[]string `json:"global_eip_id,omitempty"`

	// 全局弹性IP(有掩码)的ID
	GlobalEipSegmentId *[]string `json:"global_eip_segment_id,omitempty"`
}

func (o ListGlobalEipsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalEipsRequest struct{}"
	}

	return strings.Join([]string{"ListGlobalEipsRequest", string(data)}, " ")
}

type ListGlobalEipsRequestSortDir struct {
	value string
}

type ListGlobalEipsRequestSortDirEnum struct {
	ASC  ListGlobalEipsRequestSortDir
	DESC ListGlobalEipsRequestSortDir
}

func GetListGlobalEipsRequestSortDirEnum() ListGlobalEipsRequestSortDirEnum {
	return ListGlobalEipsRequestSortDirEnum{
		ASC: ListGlobalEipsRequestSortDir{
			value: "asc",
		},
		DESC: ListGlobalEipsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListGlobalEipsRequestSortDir) Value() string {
	return c.value
}

func (c ListGlobalEipsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGlobalEipsRequestSortDir) UnmarshalJSON(b []byte) error {
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
