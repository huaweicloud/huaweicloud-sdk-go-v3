package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CollectColumnsResponse Response Object
type CollectColumnsResponse struct {

	// 查询结果id
	Id *string `json:"id,omitempty"`

	// 查询状态 pending：处理中 failed：失败 success：成功
	Status *CollectColumnsResponseStatus `json:"status,omitempty"`

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CollectColumnsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectColumnsResponse struct{}"
	}

	return strings.Join([]string{"CollectColumnsResponse", string(data)}, " ")
}

type CollectColumnsResponseStatus struct {
	value string
}

type CollectColumnsResponseStatusEnum struct {
	PENDING CollectColumnsResponseStatus
	FAILED  CollectColumnsResponseStatus
	SUCCESS CollectColumnsResponseStatus
}

func GetCollectColumnsResponseStatusEnum() CollectColumnsResponseStatusEnum {
	return CollectColumnsResponseStatusEnum{
		PENDING: CollectColumnsResponseStatus{
			value: "pending",
		},
		FAILED: CollectColumnsResponseStatus{
			value: "failed",
		},
		SUCCESS: CollectColumnsResponseStatus{
			value: "success",
		},
	}
}

func (c CollectColumnsResponseStatus) Value() string {
	return c.value
}

func (c CollectColumnsResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CollectColumnsResponseStatus) UnmarshalJSON(b []byte) error {
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
