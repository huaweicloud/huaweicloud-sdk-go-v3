package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CollectPositionAsyncResponse Response Object
type CollectPositionAsyncResponse struct {

	// 查询结果id
	Id *string `json:"id,omitempty"`

	// 查询状态
	Status *CollectPositionAsyncResponseStatus `json:"status,omitempty"`

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CollectPositionAsyncResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectPositionAsyncResponse struct{}"
	}

	return strings.Join([]string{"CollectPositionAsyncResponse", string(data)}, " ")
}

type CollectPositionAsyncResponseStatus struct {
	value string
}

type CollectPositionAsyncResponseStatusEnum struct {
	PENDING CollectPositionAsyncResponseStatus
	FAILED  CollectPositionAsyncResponseStatus
	SUCCESS CollectPositionAsyncResponseStatus
}

func GetCollectPositionAsyncResponseStatusEnum() CollectPositionAsyncResponseStatusEnum {
	return CollectPositionAsyncResponseStatusEnum{
		PENDING: CollectPositionAsyncResponseStatus{
			value: "pending",
		},
		FAILED: CollectPositionAsyncResponseStatus{
			value: "failed",
		},
		SUCCESS: CollectPositionAsyncResponseStatus{
			value: "success",
		},
	}
}

func (c CollectPositionAsyncResponseStatus) Value() string {
	return c.value
}

func (c CollectPositionAsyncResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CollectPositionAsyncResponseStatus) UnmarshalJSON(b []byte) error {
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
