package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListExportTasksRequest Request Object
type ListExportTasksRequest struct {

	// 文件名。
	FileName *string `json:"file_name,omitempty"`

	// 任务id。
	TaskId *string `json:"task_id,omitempty"`

	// 导出任务的状态，取值为 CREATING, SUCCESS, FAIL, EXPIRED; CREATING为进行中，SUCCESS为成功，FAIL为失败，EXPIRED为过期。
	Status *ListExportTasksRequestStatus `json:"status,omitempty"`

	// 是否已下载，取值范围：true和false，默认值false。
	IsDownload *bool `json:"is_download,omitempty"`

	// 排序字段名称，需要结合sort_type字段一起使用。 - create_time 创建时间。
	SortField *ListExportTasksRequestSortField `json:"sort_field,omitempty"`

	// 排序类型，默认升序，需要结合sort_field字段一起使用。 - ASC 升序。 - DESC 降序。
	SortType *ListExportTasksRequestSortType `json:"sort_type,omitempty"`

	// 分页偏移量，默认值：0。
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小，取值范围1-100，默认值:20。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListExportTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExportTasksRequest struct{}"
	}

	return strings.Join([]string{"ListExportTasksRequest", string(data)}, " ")
}

type ListExportTasksRequestStatus struct {
	value string
}

type ListExportTasksRequestStatusEnum struct {
	CREATING ListExportTasksRequestStatus
	SUCCESS  ListExportTasksRequestStatus
	FAIL     ListExportTasksRequestStatus
	EXPIRED  ListExportTasksRequestStatus
}

func GetListExportTasksRequestStatusEnum() ListExportTasksRequestStatusEnum {
	return ListExportTasksRequestStatusEnum{
		CREATING: ListExportTasksRequestStatus{
			value: "CREATING",
		},
		SUCCESS: ListExportTasksRequestStatus{
			value: "SUCCESS",
		},
		FAIL: ListExportTasksRequestStatus{
			value: "FAIL",
		},
		EXPIRED: ListExportTasksRequestStatus{
			value: "EXPIRED",
		},
	}
}

func (c ListExportTasksRequestStatus) Value() string {
	return c.value
}

func (c ListExportTasksRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListExportTasksRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListExportTasksRequestSortField struct {
	value string
}

type ListExportTasksRequestSortFieldEnum struct {
	CREATE_TIME ListExportTasksRequestSortField
}

func GetListExportTasksRequestSortFieldEnum() ListExportTasksRequestSortFieldEnum {
	return ListExportTasksRequestSortFieldEnum{
		CREATE_TIME: ListExportTasksRequestSortField{
			value: "create_time",
		},
	}
}

func (c ListExportTasksRequestSortField) Value() string {
	return c.value
}

func (c ListExportTasksRequestSortField) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListExportTasksRequestSortField) UnmarshalJSON(b []byte) error {
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

type ListExportTasksRequestSortType struct {
	value string
}

type ListExportTasksRequestSortTypeEnum struct {
	ASC  ListExportTasksRequestSortType
	DESC ListExportTasksRequestSortType
}

func GetListExportTasksRequestSortTypeEnum() ListExportTasksRequestSortTypeEnum {
	return ListExportTasksRequestSortTypeEnum{
		ASC: ListExportTasksRequestSortType{
			value: "ASC",
		},
		DESC: ListExportTasksRequestSortType{
			value: "DESC",
		},
	}
}

func (c ListExportTasksRequestSortType) Value() string {
	return c.value
}

func (c ListExportTasksRequestSortType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListExportTasksRequestSortType) UnmarshalJSON(b []byte) error {
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
