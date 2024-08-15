package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GetScriptJobBatchRequest Request Object
type GetScriptJobBatchRequest struct {

	// 批次index
	BatchIndex int32 `json:"batch_index"`

	// 脚本工单的执行Id，取自executeJobScript和ListJobScriptOrders返回体中
	ExecuteUuid string `json:"execute_uuid"`

	// 实例执行状态 READY：待执行 PROCESSING：执行中 ABNORMAL：异常 CANCELED：已取消 FINISHED：成功
	Status *GetScriptJobBatchRequestStatus `json:"status,omitempty"`

	// 分页参数：每页返回记录个数限制
	Limit int32 `json:"limit"`

	// 分页参数：上一页最后一个记录id
	Marker int64 `json:"marker"`

	// 国际化标记，zh-cn表示中文，en-us或不传表示英文
	XLanguage *string `json:"X-Language,omitempty"`

	// 项目ID，一个项目对应一个region
	XProjectId *string `json:"x-project-id,omitempty"`

	// IAM5.0用户信息
	XUserProfile *string `json:"x-user-profile,omitempty"`
}

func (o GetScriptJobBatchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetScriptJobBatchRequest struct{}"
	}

	return strings.Join([]string{"GetScriptJobBatchRequest", string(data)}, " ")
}

type GetScriptJobBatchRequestStatus struct {
	value string
}

type GetScriptJobBatchRequestStatusEnum struct {
	READY      GetScriptJobBatchRequestStatus
	PROCESSING GetScriptJobBatchRequestStatus
	ABNORMAL   GetScriptJobBatchRequestStatus
	CANCELED   GetScriptJobBatchRequestStatus
	FINISHED   GetScriptJobBatchRequestStatus
}

func GetGetScriptJobBatchRequestStatusEnum() GetScriptJobBatchRequestStatusEnum {
	return GetScriptJobBatchRequestStatusEnum{
		READY: GetScriptJobBatchRequestStatus{
			value: "READY",
		},
		PROCESSING: GetScriptJobBatchRequestStatus{
			value: "PROCESSING",
		},
		ABNORMAL: GetScriptJobBatchRequestStatus{
			value: "ABNORMAL",
		},
		CANCELED: GetScriptJobBatchRequestStatus{
			value: "CANCELED",
		},
		FINISHED: GetScriptJobBatchRequestStatus{
			value: "FINISHED",
		},
	}
}

func (c GetScriptJobBatchRequestStatus) Value() string {
	return c.value
}

func (c GetScriptJobBatchRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetScriptJobBatchRequestStatus) UnmarshalJSON(b []byte) error {
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
