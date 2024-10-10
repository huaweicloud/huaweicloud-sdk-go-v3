package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListJobDdlsRequest Request Object
type ListJobDdlsRequest struct {

	// 请求语言类型。
	XLanguage *string `json:"X-Language,omitempty"`

	// 偏移量，默认值为0，表示查询该偏移量后面的记录。
	Offset *int32 `json:"offset,omitempty"`

	// 查询返回记录的数量限制.默认值为10。
	Limit *int32 `json:"limit,omitempty"`

	// DDL序列起始值。
	StartSeqNo *int64 `json:"start_seq_no,omitempty"`

	// DDL序列结束值。
	EndSeqNo *int64 `json:"end_seq_no,omitempty"`

	// DDL状态，0为无告警，1为告警中。
	Status *ListJobDdlsRequestStatus `json:"status,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`
}

func (o ListJobDdlsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobDdlsRequest struct{}"
	}

	return strings.Join([]string{"ListJobDdlsRequest", string(data)}, " ")
}

type ListJobDdlsRequestStatus struct {
	value int32
}

type ListJobDdlsRequestStatusEnum struct {
	E_0 ListJobDdlsRequestStatus
	E_1 ListJobDdlsRequestStatus
}

func GetListJobDdlsRequestStatusEnum() ListJobDdlsRequestStatusEnum {
	return ListJobDdlsRequestStatusEnum{
		E_0: ListJobDdlsRequestStatus{
			value: 0,
		}, E_1: ListJobDdlsRequestStatus{
			value: 1,
		},
	}
}

func (c ListJobDdlsRequestStatus) Value() int32 {
	return c.value
}

func (c ListJobDdlsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListJobDdlsRequestStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
