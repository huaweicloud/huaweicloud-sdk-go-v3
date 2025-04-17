package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DdosConfig WeeklyReport
type DdosConfig struct {

	// 是否开启L7层防护，固定值为fasle
	EnableL7 bool `json:"enable_L7"`

	// 流量分段ID，取值范围：1：10M;2：30M;3：50M;4：70M;5：100M;6：150M;7：200M;8：250M;9：300M;10：500M;11：800M;88：1000M;99：默认防护。
	TrafficPosId int64 `json:"traffic_pos_id"`

	// HTTP请求数分段ID，固定值为1
	HttpRequestPosId int64 `json:"http_request_pos_id"`

	// 清洗时访问限制分段ID，取值范围：1：10M;2：30M;3：50M;4：70M;5：100M;6：150M;7：200M;8：250M;9：300M;10：500M;11：800M;88：1000M;99：默认防护。
	CleaningAccessPosId int64 `json:"cleaning_access_pos_id"`

	// 应用类型ID，固定值为0
	AppTypeId DdosConfigAppTypeId `json:"app_type_id"`
}

func (o DdosConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DdosConfig struct{}"
	}

	return strings.Join([]string{"DdosConfig", string(data)}, " ")
}

type DdosConfigAppTypeId struct {
	value int64
}

type DdosConfigAppTypeIdEnum struct {
	E_0 DdosConfigAppTypeId
	E_1 DdosConfigAppTypeId
}

func GetDdosConfigAppTypeIdEnum() DdosConfigAppTypeIdEnum {
	return DdosConfigAppTypeIdEnum{
		E_0: DdosConfigAppTypeId{
			value: 0,
		}, E_1: DdosConfigAppTypeId{
			value: 1,
		},
	}
}

func (c DdosConfigAppTypeId) Value() int64 {
	return c.value
}

func (c DdosConfigAppTypeId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DdosConfigAppTypeId) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int64")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int64")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int64); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int64 error")
	}
}
