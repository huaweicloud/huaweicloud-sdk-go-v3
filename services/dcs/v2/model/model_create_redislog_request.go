package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateRedislogRequest Request Object
type CreateRedislogRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 日期偏移量，表示从过去的n天开始查询，例如：传入0则表示查询今天的日志，传入7则表示查询过去7天的日志。当前支持0，1，3，7。
	QueryTime *CreateRedislogRequestQueryTime `json:"query_time,omitempty"`

	// 返回日志的类型，当前仅支持Redis运行日志，类型为run
	LogType string `json:"log_type"`

	// 副本ID，可以从分片与副本中查询对应节点的副本ID
	ReplicationId *string `json:"replication_id,omitempty"`
}

func (o CreateRedislogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRedislogRequest struct{}"
	}

	return strings.Join([]string{"CreateRedislogRequest", string(data)}, " ")
}

type CreateRedislogRequestQueryTime struct {
	value int32
}

type CreateRedislogRequestQueryTimeEnum struct {
	E_0 CreateRedislogRequestQueryTime
	E_1 CreateRedislogRequestQueryTime
	E_3 CreateRedislogRequestQueryTime
	E_7 CreateRedislogRequestQueryTime
}

func GetCreateRedislogRequestQueryTimeEnum() CreateRedislogRequestQueryTimeEnum {
	return CreateRedislogRequestQueryTimeEnum{
		E_0: CreateRedislogRequestQueryTime{
			value: 0,
		}, E_1: CreateRedislogRequestQueryTime{
			value: 1,
		}, E_3: CreateRedislogRequestQueryTime{
			value: 3,
		}, E_7: CreateRedislogRequestQueryTime{
			value: 7,
		},
	}
}

func (c CreateRedislogRequestQueryTime) Value() int32 {
	return c.value
}

func (c CreateRedislogRequestQueryTime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRedislogRequestQueryTime) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
