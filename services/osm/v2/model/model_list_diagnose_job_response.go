package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ListDiagnoseJobResponse struct {

	//  任务执行的状态 0：准备中，2：执行中，3：完成，4：失败，7：未执行，8：不可用；用于判断任务的是否执行结束，3就是结束了，4,7,8说明是TSC诊断脚本有问题，OSM这边无法处理
	Status *ListDiagnoseJobResponseStatus `json:"status,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 任务ID
	JobId *int32 `json:"job_id,omitempty"`

	// 租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 任务的检查项结果
	ItemsResult *[]ItemResultVo `json:"items_result,omitempty"`

	// 任务创建时间
	CreateTime     *int64 `json:"create_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDiagnoseJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDiagnoseJobResponse struct{}"
	}

	return strings.Join([]string{"ListDiagnoseJobResponse", string(data)}, " ")
}

type ListDiagnoseJobResponseStatus struct {
	value int32
}

type ListDiagnoseJobResponseStatusEnum struct {
	E_0 ListDiagnoseJobResponseStatus
	E_2 ListDiagnoseJobResponseStatus
	E_3 ListDiagnoseJobResponseStatus
	E_4 ListDiagnoseJobResponseStatus
	E_7 ListDiagnoseJobResponseStatus
	E_8 ListDiagnoseJobResponseStatus
}

func GetListDiagnoseJobResponseStatusEnum() ListDiagnoseJobResponseStatusEnum {
	return ListDiagnoseJobResponseStatusEnum{
		E_0: ListDiagnoseJobResponseStatus{
			value: 0,
		}, E_2: ListDiagnoseJobResponseStatus{
			value: 2,
		}, E_3: ListDiagnoseJobResponseStatus{
			value: 3,
		}, E_4: ListDiagnoseJobResponseStatus{
			value: 4,
		}, E_7: ListDiagnoseJobResponseStatus{
			value: 7,
		}, E_8: ListDiagnoseJobResponseStatus{
			value: 8,
		},
	}
}

func (c ListDiagnoseJobResponseStatus) Value() int32 {
	return c.value
}

func (c ListDiagnoseJobResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDiagnoseJobResponseStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
