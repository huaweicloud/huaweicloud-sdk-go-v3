package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListJobsRequest Request Object
type ListJobsRequest struct {

	// 每页条数
	Limit *int32 `json:"limit,omitempty"`

	// 分页起始点
	Offset *int32 `json:"offset,omitempty"`

	// 分页起始点
	Marker *string `json:"marker,omitempty"`

	// 翻页方向
	PageReverse *bool `json:"page_reverse,omitempty"`

	Fields *[]ListJobsRequestFields `json:"fields,omitempty"`

	// 按照sort_key指定的字段排序
	SortKey *[]ListJobsRequestSortKey `json:"sort_key,omitempty"`

	// 排序的方向，倒序或者正序
	SortDir *[]ListJobsRequestSortDir `json:"sort_dir,omitempty"`

	Id *[]string `json:"id,omitempty"`

	Action *[]ListJobsRequestAction `json:"action,omitempty"`

	Status *[]ListJobsRequestStatus `json:"status,omitempty"`
}

func (o ListJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobsRequest struct{}"
	}

	return strings.Join([]string{"ListJobsRequest", string(data)}, " ")
}

type ListJobsRequestFields struct {
	value string
}

type ListJobsRequestFieldsEnum struct {
	ID            ListJobsRequestFields
	ACTION        ListJobsRequestFields
	STATUS        ListJobsRequestFields
	ERROR_TASK    ListJobsRequestFields
	ERROR_CODE    ListJobsRequestFields
	ERROR_MESSAGE ListJobsRequestFields
	START_TIME    ListJobsRequestFields
	END_TIME      ListJobsRequestFields
}

func GetListJobsRequestFieldsEnum() ListJobsRequestFieldsEnum {
	return ListJobsRequestFieldsEnum{
		ID: ListJobsRequestFields{
			value: "id",
		},
		ACTION: ListJobsRequestFields{
			value: "action",
		},
		STATUS: ListJobsRequestFields{
			value: "status",
		},
		ERROR_TASK: ListJobsRequestFields{
			value: "error_task",
		},
		ERROR_CODE: ListJobsRequestFields{
			value: "error_code",
		},
		ERROR_MESSAGE: ListJobsRequestFields{
			value: "error_message",
		},
		START_TIME: ListJobsRequestFields{
			value: "start_time",
		},
		END_TIME: ListJobsRequestFields{
			value: "end_time",
		},
	}
}

func (c ListJobsRequestFields) Value() string {
	return c.value
}

func (c ListJobsRequestFields) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListJobsRequestFields) UnmarshalJSON(b []byte) error {
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

type ListJobsRequestSortKey struct {
	value string
}

type ListJobsRequestSortKeyEnum struct {
	ID         ListJobsRequestSortKey
	START_TIME ListJobsRequestSortKey
	END_TIME   ListJobsRequestSortKey
}

func GetListJobsRequestSortKeyEnum() ListJobsRequestSortKeyEnum {
	return ListJobsRequestSortKeyEnum{
		ID: ListJobsRequestSortKey{
			value: "id",
		},
		START_TIME: ListJobsRequestSortKey{
			value: "start_time",
		},
		END_TIME: ListJobsRequestSortKey{
			value: "end_time",
		},
	}
}

func (c ListJobsRequestSortKey) Value() string {
	return c.value
}

func (c ListJobsRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListJobsRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListJobsRequestSortDir struct {
	value string
}

type ListJobsRequestSortDirEnum struct {
	ASC  ListJobsRequestSortDir
	DESC ListJobsRequestSortDir
}

func GetListJobsRequestSortDirEnum() ListJobsRequestSortDirEnum {
	return ListJobsRequestSortDirEnum{
		ASC: ListJobsRequestSortDir{
			value: "asc",
		},
		DESC: ListJobsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListJobsRequestSortDir) Value() string {
	return c.value
}

func (c ListJobsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListJobsRequestSortDir) UnmarshalJSON(b []byte) error {
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

type ListJobsRequestAction struct {
	value string
}

type ListJobsRequestActionEnum struct {
	CREATE_GEIP_GRAPH            ListJobsRequestAction
	UPDATE_GEIP_GRAPH            ListJobsRequestAction
	BIND_GEIP_GRAPH              ListJobsRequestAction
	UN_BIND_GEIP_GRAPH           ListJobsRequestAction
	BATCH_ATTACH_BANDWIDTH_GRAPH ListJobsRequestAction
	BATCH_DETACH_BANDWIDTH_GRAPH ListJobsRequestAction
	CREATE_GEIP_SEGMENT_GRAPH    ListJobsRequestAction
	BIND_GEIP_SEGMENT_GRAPH      ListJobsRequestAction
}

func GetListJobsRequestActionEnum() ListJobsRequestActionEnum {
	return ListJobsRequestActionEnum{
		CREATE_GEIP_GRAPH: ListJobsRequestAction{
			value: "CreateGEIPGraph",
		},
		UPDATE_GEIP_GRAPH: ListJobsRequestAction{
			value: "UpdateGEIPGraph",
		},
		BIND_GEIP_GRAPH: ListJobsRequestAction{
			value: "BindGEIPGraph",
		},
		UN_BIND_GEIP_GRAPH: ListJobsRequestAction{
			value: "UnBindGEIPGraph",
		},
		BATCH_ATTACH_BANDWIDTH_GRAPH: ListJobsRequestAction{
			value: "BatchAttachBandwidthGraph",
		},
		BATCH_DETACH_BANDWIDTH_GRAPH: ListJobsRequestAction{
			value: "BatchDetachBandwidthGraph",
		},
		CREATE_GEIP_SEGMENT_GRAPH: ListJobsRequestAction{
			value: "CreateGeipSegmentGraph",
		},
		BIND_GEIP_SEGMENT_GRAPH: ListJobsRequestAction{
			value: "BindGeipSegmentGraph",
		},
	}
}

func (c ListJobsRequestAction) Value() string {
	return c.value
}

func (c ListJobsRequestAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListJobsRequestAction) UnmarshalJSON(b []byte) error {
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

type ListJobsRequestStatus struct {
	value string
}

type ListJobsRequestStatusEnum struct {
	FINISH_SUCC          ListJobsRequestStatus
	FINISH_ROLLBACK_SUCC ListJobsRequestStatus
}

func GetListJobsRequestStatusEnum() ListJobsRequestStatusEnum {
	return ListJobsRequestStatusEnum{
		FINISH_SUCC: ListJobsRequestStatus{
			value: "FINISH_SUCC",
		},
		FINISH_ROLLBACK_SUCC: ListJobsRequestStatus{
			value: "FINISH_ROLLBACK_SUCC",
		},
	}
}

func (c ListJobsRequestStatus) Value() string {
	return c.value
}

func (c ListJobsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListJobsRequestStatus) UnmarshalJSON(b []byte) error {
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
