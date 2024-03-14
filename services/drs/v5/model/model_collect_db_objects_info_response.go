package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CollectDbObjectsInfoResponse Response Object
type CollectDbObjectsInfoResponse struct {

	// 查询结果id
	Id *string `json:"id,omitempty"`

	// 查询状态
	Status *CollectDbObjectsInfoResponseStatus `json:"status,omitempty"`

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CollectDbObjectsInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectDbObjectsInfoResponse struct{}"
	}

	return strings.Join([]string{"CollectDbObjectsInfoResponse", string(data)}, " ")
}

type CollectDbObjectsInfoResponseStatus struct {
	value string
}

type CollectDbObjectsInfoResponseStatusEnum struct {
	PENDING CollectDbObjectsInfoResponseStatus
	FAILED  CollectDbObjectsInfoResponseStatus
	SUCCESS CollectDbObjectsInfoResponseStatus
}

func GetCollectDbObjectsInfoResponseStatusEnum() CollectDbObjectsInfoResponseStatusEnum {
	return CollectDbObjectsInfoResponseStatusEnum{
		PENDING: CollectDbObjectsInfoResponseStatus{
			value: "pending",
		},
		FAILED: CollectDbObjectsInfoResponseStatus{
			value: "failed",
		},
		SUCCESS: CollectDbObjectsInfoResponseStatus{
			value: "success",
		},
	}
}

func (c CollectDbObjectsInfoResponseStatus) Value() string {
	return c.value
}

func (c CollectDbObjectsInfoResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CollectDbObjectsInfoResponseStatus) UnmarshalJSON(b []byte) error {
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
