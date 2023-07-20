package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListFlowLogsRequest Request Object
type ListFlowLogsRequest struct {

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	// 采集的资源类型
	ResourceType *ListFlowLogsRequestResourceType `json:"resource_type,omitempty"`

	// 连接对应的资源ID列表
	ResourceId *[]string `json:"resource_id,omitempty"`

	// 每页返回的个数。 取值范围：0~2000。
	Limit *int32 `json:"limit,omitempty"`

	// 上一页最后一条记录的企业路由器实例的id，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 按关键字排序，默认按照id排序，可选值:id|name|state
	SortKey *[]string `json:"sort_key,omitempty"`

	// 返回结果按照升序或降序排列，默认为asc,降序为desc
	SortDir *[]ListFlowLogsRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListFlowLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlowLogsRequest struct{}"
	}

	return strings.Join([]string{"ListFlowLogsRequest", string(data)}, " ")
}

type ListFlowLogsRequestResourceType struct {
	value string
}

type ListFlowLogsRequestResourceTypeEnum struct {
	ATTACHMENT ListFlowLogsRequestResourceType
}

func GetListFlowLogsRequestResourceTypeEnum() ListFlowLogsRequestResourceTypeEnum {
	return ListFlowLogsRequestResourceTypeEnum{
		ATTACHMENT: ListFlowLogsRequestResourceType{
			value: "attachment",
		},
	}
}

func (c ListFlowLogsRequestResourceType) Value() string {
	return c.value
}

func (c ListFlowLogsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlowLogsRequestResourceType) UnmarshalJSON(b []byte) error {
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

type ListFlowLogsRequestSortDir struct {
	value string
}

type ListFlowLogsRequestSortDirEnum struct {
	ASC  ListFlowLogsRequestSortDir
	DESC ListFlowLogsRequestSortDir
}

func GetListFlowLogsRequestSortDirEnum() ListFlowLogsRequestSortDirEnum {
	return ListFlowLogsRequestSortDirEnum{
		ASC: ListFlowLogsRequestSortDir{
			value: "asc",
		},
		DESC: ListFlowLogsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListFlowLogsRequestSortDir) Value() string {
	return c.value
}

func (c ListFlowLogsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlowLogsRequestSortDir) UnmarshalJSON(b []byte) error {
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
