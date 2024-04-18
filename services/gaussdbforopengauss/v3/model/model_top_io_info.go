package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TopIoInfo Top IO信息
type TopIoInfo struct {

	// 线程ID
	ThreadId *string `json:"thread_id,omitempty"`

	// 线程分类标识，取值：业务（worker）和后台（background）。需将GUC参数\"enable_thread_pool\"设置为on
	ThreadType *TopIoInfoThreadType `json:"thread_type,omitempty"`

	// 从磁盘读取数据速率, 单位：KB/s
	DiskReadRate *int32 `json:"disk_read_rate,omitempty"`

	// 写入磁盘数据速率, 单位：KB/s
	DiskWriteRate *int32 `json:"disk_write_rate,omitempty"`

	// 会话ID
	SessionId *string `json:"session_id,omitempty"`

	// SQL ID
	UniqueSqlId *string `json:"unique_sql_id,omitempty"`

	// 数据库
	DatabaseName *string `json:"database_name,omitempty"`

	// 客户端IP
	ClientIp *string `json:"client_ip,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 状态
	State *string `json:"state,omitempty"`

	// 语句开始时间
	SqlStart *int64 `json:"sql_start,omitempty"`
}

func (o TopIoInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopIoInfo struct{}"
	}

	return strings.Join([]string{"TopIoInfo", string(data)}, " ")
}

type TopIoInfoThreadType struct {
	value string
}

type TopIoInfoThreadTypeEnum struct {
	WORKER     TopIoInfoThreadType
	BACKGROUND TopIoInfoThreadType
}

func GetTopIoInfoThreadTypeEnum() TopIoInfoThreadTypeEnum {
	return TopIoInfoThreadTypeEnum{
		WORKER: TopIoInfoThreadType{
			value: "worker",
		},
		BACKGROUND: TopIoInfoThreadType{
			value: "background",
		},
	}
}

func (c TopIoInfoThreadType) Value() string {
	return c.value
}

func (c TopIoInfoThreadType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TopIoInfoThreadType) UnmarshalJSON(b []byte) error {
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
