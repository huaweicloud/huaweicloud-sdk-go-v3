package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CollectDbObjectsAsyncResponse Response Object
type CollectDbObjectsAsyncResponse struct {

	// 查询结果id
	Id *string `json:"id,omitempty"`

	// 查询状态
	Status *CollectDbObjectsAsyncResponseStatus `json:"status,omitempty"`

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CollectDbObjectsAsyncResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectDbObjectsAsyncResponse struct{}"
	}

	return strings.Join([]string{"CollectDbObjectsAsyncResponse", string(data)}, " ")
}

type CollectDbObjectsAsyncResponseStatus struct {
	value string
}

type CollectDbObjectsAsyncResponseStatusEnum struct {
	PENDING CollectDbObjectsAsyncResponseStatus
	FAILED  CollectDbObjectsAsyncResponseStatus
	SUCCESS CollectDbObjectsAsyncResponseStatus
}

func GetCollectDbObjectsAsyncResponseStatusEnum() CollectDbObjectsAsyncResponseStatusEnum {
	return CollectDbObjectsAsyncResponseStatusEnum{
		PENDING: CollectDbObjectsAsyncResponseStatus{
			value: "pending",
		},
		FAILED: CollectDbObjectsAsyncResponseStatus{
			value: "failed",
		},
		SUCCESS: CollectDbObjectsAsyncResponseStatus{
			value: "success",
		},
	}
}

func (c CollectDbObjectsAsyncResponseStatus) Value() string {
	return c.value
}

func (c CollectDbObjectsAsyncResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CollectDbObjectsAsyncResponseStatus) UnmarshalJSON(b []byte) error {
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
