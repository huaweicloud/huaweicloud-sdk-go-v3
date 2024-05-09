package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchJobActionReq 任务动作请求体。
type BatchJobActionReq struct {

	// 需要执行的特定操作。
	Action BatchJobActionReqAction `json:"action"`

	// 任务ID（集群模式 取父任务的任务ID）。
	JobId string `json:"job_id"`

	// 操作对应的参数（API参考文档-批量测试连接-集群模式-property字段数据结构说明）[字段说明参考](https://support.huaweicloud.com/api-drs/drs_03_0106.html)
	Property string `json:"property"`
}

func (o BatchJobActionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchJobActionReq struct{}"
	}

	return strings.Join([]string{"BatchJobActionReq", string(data)}, " ")
}

type BatchJobActionReqAction struct {
	value string
}

type BatchJobActionReqActionEnum struct {
	TEST_CONNECTION BatchJobActionReqAction
}

func GetBatchJobActionReqActionEnum() BatchJobActionReqActionEnum {
	return BatchJobActionReqActionEnum{
		TEST_CONNECTION: BatchJobActionReqAction{
			value: "testConnection",
		},
	}
}

func (c BatchJobActionReqAction) Value() string {
	return c.value
}

func (c BatchJobActionReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchJobActionReqAction) UnmarshalJSON(b []byte) error {
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
