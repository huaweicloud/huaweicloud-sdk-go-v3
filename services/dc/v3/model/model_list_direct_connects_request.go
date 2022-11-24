package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListDirectConnectsRequest struct {

	// 每页返回的个数。 取值范围：1~2000。
	Limit *int32 `json:"limit,omitempty"`

	// 上一页最后一条资源记录的ID，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 显示字段列表
	Fields *[]string `json:"fields,omitempty"`

	// 排序字段。
	SortKey *string `json:"sort_key,omitempty"`

	// 返回结果按照升序(asc)或降序(desc)排列，默认为asc
	SortDir *[]ListDirectConnectsRequestSortDir `json:"sort_dir,omitempty"`

	// 根椐运营专线ID过滤托管专线列表
	HostingId *[]string `json:"hosting_id,omitempty"`

	// 根据企业项目ID过滤中心网络列表。
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// 根据资源ID过滤实例
	Id *[]string `json:"id,omitempty"`

	// 根据名字过滤查询，可查询多个名字。
	Name *[]string `json:"name,omitempty"`
}

func (o ListDirectConnectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDirectConnectsRequest struct{}"
	}

	return strings.Join([]string{"ListDirectConnectsRequest", string(data)}, " ")
}

type ListDirectConnectsRequestSortDir struct {
	value string
}

type ListDirectConnectsRequestSortDirEnum struct {
	ASC  ListDirectConnectsRequestSortDir
	DESC ListDirectConnectsRequestSortDir
}

func GetListDirectConnectsRequestSortDirEnum() ListDirectConnectsRequestSortDirEnum {
	return ListDirectConnectsRequestSortDirEnum{
		ASC: ListDirectConnectsRequestSortDir{
			value: "asc",
		},
		DESC: ListDirectConnectsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListDirectConnectsRequestSortDir) Value() string {
	return c.value
}

func (c ListDirectConnectsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDirectConnectsRequestSortDir) UnmarshalJSON(b []byte) error {
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
