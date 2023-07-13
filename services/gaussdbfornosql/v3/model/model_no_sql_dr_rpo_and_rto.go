package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NoSqlDrRpoAndRto 容灾实例切换的RPO和RTO指标
type NoSqlDrRpoAndRto struct {

	// 场景，枚举值   FAILOVER 强制切换;    SWITCHOVER 主备倒换
	Scene NoSqlDrRpoAndRtoScene `json:"scene"`

	// 倒换或切换丢失数据时长，单位：秒（s）
	Rpo *int64 `json:"rpo,omitempty"`

	// 倒换或切换恢复时长，单位：秒（s）
	Rto *int64 `json:"rto,omitempty"`
}

func (o NoSqlDrRpoAndRto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NoSqlDrRpoAndRto struct{}"
	}

	return strings.Join([]string{"NoSqlDrRpoAndRto", string(data)}, " ")
}

type NoSqlDrRpoAndRtoScene struct {
	value string
}

type NoSqlDrRpoAndRtoSceneEnum struct {
	FAILOVER   NoSqlDrRpoAndRtoScene
	SWITCHOVER NoSqlDrRpoAndRtoScene
}

func GetNoSqlDrRpoAndRtoSceneEnum() NoSqlDrRpoAndRtoSceneEnum {
	return NoSqlDrRpoAndRtoSceneEnum{
		FAILOVER: NoSqlDrRpoAndRtoScene{
			value: "FAILOVER",
		},
		SWITCHOVER: NoSqlDrRpoAndRtoScene{
			value: "SWITCHOVER",
		},
	}
}

func (c NoSqlDrRpoAndRtoScene) Value() string {
	return c.value
}

func (c NoSqlDrRpoAndRtoScene) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NoSqlDrRpoAndRtoScene) UnmarshalJSON(b []byte) error {
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
