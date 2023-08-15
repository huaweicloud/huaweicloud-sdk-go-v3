package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListHostedDirectConnectsRequest Request Object
type ListHostedDirectConnectsRequest struct {

	// 每页返回的个数。 取值范围：1~2000。
	Limit *int32 `json:"limit,omitempty"`

	// 上一页最后一条资源记录的ID，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 显示字段列表
	Fields *[]string `json:"fields,omitempty"`

	// 返回结果按照升序(asc)或降序(desc)排列，默认为asc
	SortDir *[]ListHostedDirectConnectsRequestSortDir `json:"sort_dir,omitempty"`

	// 排序字段。
	SortKey *string `json:"sort_key,omitempty"`

	// 根椐运营专线ID过滤托管专线列表
	HostingId *[]string `json:"hosting_id,omitempty"`

	// 根据资源ID过滤实例
	Id *[]string `json:"id,omitempty"`

	// 根据名字过滤查询，可查询多个名字。
	Name *[]string `json:"name,omitempty"`
}

func (o ListHostedDirectConnectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostedDirectConnectsRequest struct{}"
	}

	return strings.Join([]string{"ListHostedDirectConnectsRequest", string(data)}, " ")
}

type ListHostedDirectConnectsRequestSortDir struct {
	value string
}

type ListHostedDirectConnectsRequestSortDirEnum struct {
	ASC  ListHostedDirectConnectsRequestSortDir
	DESC ListHostedDirectConnectsRequestSortDir
}

func GetListHostedDirectConnectsRequestSortDirEnum() ListHostedDirectConnectsRequestSortDirEnum {
	return ListHostedDirectConnectsRequestSortDirEnum{
		ASC: ListHostedDirectConnectsRequestSortDir{
			value: "asc",
		},
		DESC: ListHostedDirectConnectsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListHostedDirectConnectsRequestSortDir) Value() string {
	return c.value
}

func (c ListHostedDirectConnectsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHostedDirectConnectsRequestSortDir) UnmarshalJSON(b []byte) error {
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
