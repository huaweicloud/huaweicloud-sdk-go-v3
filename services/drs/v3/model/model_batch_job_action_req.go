package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 任务动作请求体
type BatchJobActionReq struct {
	// 执行操作

	Action BatchJobActionReqAction `json:"action"`
	// 任务ID（集群模式 取父任务的任务id）

	JobId string `json:"job_id"`
	// 操作对应的参数（API参考文档-批量测试连接-集群模式-property字段数据结构说明）[字段说明参考](https://support.huaweicloud.com/api-drs/zh-cn_topic_0295171516.html)

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

func (c BatchJobActionReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchJobActionReqAction) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
