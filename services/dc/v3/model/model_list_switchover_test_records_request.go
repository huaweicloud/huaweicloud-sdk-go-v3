package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSwitchoverTestRecordsRequest Request Object
type ListSwitchoverTestRecordsRequest struct {

	// 每页返回的个数。 取值范围：1~2000。
	Limit *int32 `json:"limit,omitempty"`

	// 上一页最后一条资源记录的ID，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 显示字段列表
	Fields *[]string `json:"fields,omitempty"`

	// 返回结果按照升序(asc)或降序(desc)排列，默认为asc
	SortDir *[]ListSwitchoverTestRecordsRequestSortDir `json:"sort_dir,omitempty"`

	// 排序字段。
	SortKey *string `json:"sort_key,omitempty"`

	// 通过RESOURCE-ID过虑倒换测试记录信息
	ResourceId *[]string `json:"resource_id,omitempty"`
}

func (o ListSwitchoverTestRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSwitchoverTestRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListSwitchoverTestRecordsRequest", string(data)}, " ")
}

type ListSwitchoverTestRecordsRequestSortDir struct {
	value string
}

type ListSwitchoverTestRecordsRequestSortDirEnum struct {
	ASC  ListSwitchoverTestRecordsRequestSortDir
	DESC ListSwitchoverTestRecordsRequestSortDir
}

func GetListSwitchoverTestRecordsRequestSortDirEnum() ListSwitchoverTestRecordsRequestSortDirEnum {
	return ListSwitchoverTestRecordsRequestSortDirEnum{
		ASC: ListSwitchoverTestRecordsRequestSortDir{
			value: "asc",
		},
		DESC: ListSwitchoverTestRecordsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListSwitchoverTestRecordsRequestSortDir) Value() string {
	return c.value
}

func (c ListSwitchoverTestRecordsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSwitchoverTestRecordsRequestSortDir) UnmarshalJSON(b []byte) error {
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
