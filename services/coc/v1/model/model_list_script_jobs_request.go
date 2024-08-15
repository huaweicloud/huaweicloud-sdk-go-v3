package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListScriptJobsRequest Request Object
type ListScriptJobsRequest struct {

	// 分页参数
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数
	Marker *int64 `json:"marker,omitempty"`

	// 创建时间开始
	StartTime *int64 `json:"start_time,omitempty"`

	// 创建时间结束
	EndTime *int64 `json:"end_time,omitempty"`

	// 创建人
	Creator *string `json:"creator,omitempty"`

	// 工单状态 PROCESSING：执行中 ABNORMAL：异常 PAUSED：暂停 CANCELED：已取消 FINISHED：成功
	Status *ListScriptJobsRequestStatus `json:"status,omitempty"`

	// 国际化标记，zh-cn表示中文，en-us或不传表示英文
	XLanguage *string `json:"X-Language,omitempty"`

	// 项目ID，一个项目对应一个region
	XProjectId *string `json:"x-project-id,omitempty"`

	// IAM5.0用户信息
	XUserProfile *string `json:"x-user-profile,omitempty"`
}

func (o ListScriptJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScriptJobsRequest struct{}"
	}

	return strings.Join([]string{"ListScriptJobsRequest", string(data)}, " ")
}

type ListScriptJobsRequestStatus struct {
	value string
}

type ListScriptJobsRequestStatusEnum struct {
	PROCESSING ListScriptJobsRequestStatus
	ABNORMAL   ListScriptJobsRequestStatus
	PAUSED     ListScriptJobsRequestStatus
	CANCELED   ListScriptJobsRequestStatus
	FINISHED   ListScriptJobsRequestStatus
}

func GetListScriptJobsRequestStatusEnum() ListScriptJobsRequestStatusEnum {
	return ListScriptJobsRequestStatusEnum{
		PROCESSING: ListScriptJobsRequestStatus{
			value: "PROCESSING",
		},
		ABNORMAL: ListScriptJobsRequestStatus{
			value: "ABNORMAL",
		},
		PAUSED: ListScriptJobsRequestStatus{
			value: "PAUSED",
		},
		CANCELED: ListScriptJobsRequestStatus{
			value: "CANCELED",
		},
		FINISHED: ListScriptJobsRequestStatus{
			value: "FINISHED",
		},
	}
}

func (c ListScriptJobsRequestStatus) Value() string {
	return c.value
}

func (c ListScriptJobsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListScriptJobsRequestStatus) UnmarshalJSON(b []byte) error {
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
