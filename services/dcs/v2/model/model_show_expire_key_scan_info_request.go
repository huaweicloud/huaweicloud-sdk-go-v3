package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowExpireKeyScanInfoRequest Request Object
type ShowExpireKeyScanInfoRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询， start大于等于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 过期key状态
	Status *ShowExpireKeyScanInfoRequestStatus `json:"status,omitempty"`
}

func (o ShowExpireKeyScanInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExpireKeyScanInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowExpireKeyScanInfoRequest", string(data)}, " ")
}

type ShowExpireKeyScanInfoRequestStatus struct {
	value string
}

type ShowExpireKeyScanInfoRequestStatusEnum struct {
	WAITING ShowExpireKeyScanInfoRequestStatus
	RUNNING ShowExpireKeyScanInfoRequestStatus
	SUCCESS ShowExpireKeyScanInfoRequestStatus
	FAILED  ShowExpireKeyScanInfoRequestStatus
}

func GetShowExpireKeyScanInfoRequestStatusEnum() ShowExpireKeyScanInfoRequestStatusEnum {
	return ShowExpireKeyScanInfoRequestStatusEnum{
		WAITING: ShowExpireKeyScanInfoRequestStatus{
			value: "waiting",
		},
		RUNNING: ShowExpireKeyScanInfoRequestStatus{
			value: "running",
		},
		SUCCESS: ShowExpireKeyScanInfoRequestStatus{
			value: "success",
		},
		FAILED: ShowExpireKeyScanInfoRequestStatus{
			value: "failed",
		},
	}
}

func (c ShowExpireKeyScanInfoRequestStatus) Value() string {
	return c.value
}

func (c ShowExpireKeyScanInfoRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowExpireKeyScanInfoRequestStatus) UnmarshalJSON(b []byte) error {
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
