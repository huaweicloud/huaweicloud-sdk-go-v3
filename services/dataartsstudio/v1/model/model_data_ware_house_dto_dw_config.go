package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DataWareHouseDtoDwConfig 数据开发细粒度连接配置信息
type DataWareHouseDtoDwConfig struct {

	// 是否开启细粒度认证,true表示开启细粒度认证,false表示关闭细粒度认证。
	FgacFlag *bool `json:"fgac_flag,omitempty"`

	// 细粒度认证类型，开启细粒度认证时才生效。\"0\"表示开发态细粒度认证，支持数据开发细粒度脚本运行、作业测试运行，\"1\"表示调度态细粒度认证，支持数据开发细粒度脚本运行、作业测试运行、作业执行调度。
	FgacType *string `json:"fgac_type,omitempty"`

	// 数据源连通性测试状态：   * UNKNOWN - 连通性未测试   * TESTING - 连通性测试中   * SUCCESS - 连通性测试成功   * FAILED - 连通性测试失败
	FgacConnStatus *DataWareHouseDtoDwConfigFgacConnStatus `json:"fgac_conn_status,omitempty"`

	// 最近一次连通性测试时间
	FgacConnTestTime *int64 `json:"fgac_conn_test_time,omitempty"`

	// 联通性测试失败信息，如果连通性测试成功或者未测试联通性，失败信息为空字符串
	FgacConnTestError *string `json:"fgac_conn_test_error,omitempty"`
}

func (o DataWareHouseDtoDwConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataWareHouseDtoDwConfig struct{}"
	}

	return strings.Join([]string{"DataWareHouseDtoDwConfig", string(data)}, " ")
}

type DataWareHouseDtoDwConfigFgacConnStatus struct {
	value string
}

type DataWareHouseDtoDwConfigFgacConnStatusEnum struct {
	UNKNOWN DataWareHouseDtoDwConfigFgacConnStatus
	TESTING DataWareHouseDtoDwConfigFgacConnStatus
	SUCCESS DataWareHouseDtoDwConfigFgacConnStatus
	FAILED  DataWareHouseDtoDwConfigFgacConnStatus
}

func GetDataWareHouseDtoDwConfigFgacConnStatusEnum() DataWareHouseDtoDwConfigFgacConnStatusEnum {
	return DataWareHouseDtoDwConfigFgacConnStatusEnum{
		UNKNOWN: DataWareHouseDtoDwConfigFgacConnStatus{
			value: "UNKNOWN",
		},
		TESTING: DataWareHouseDtoDwConfigFgacConnStatus{
			value: "TESTING",
		},
		SUCCESS: DataWareHouseDtoDwConfigFgacConnStatus{
			value: "SUCCESS",
		},
		FAILED: DataWareHouseDtoDwConfigFgacConnStatus{
			value: "FAILED",
		},
	}
}

func (c DataWareHouseDtoDwConfigFgacConnStatus) Value() string {
	return c.value
}

func (c DataWareHouseDtoDwConfigFgacConnStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataWareHouseDtoDwConfigFgacConnStatus) UnmarshalJSON(b []byte) error {
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
