package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type OpenAntiDDosServiceRequestBody struct {

	// 应用类型ID，固定值为0
	AppTypeId OpenAntiDDosServiceRequestBodyAppTypeId `json:"app_type_id"`

	// 清洗时访问限制分段ID，取值范围：1：10M;2：30M;3：50M;4：70M;5：100M;6：150M;7：200M;8：250M;9：300M;10：500M;11：800M;88：1000M;99：默认防护。
	CleaningAccessPosId int64 `json:"cleaning_access_pos_id"`

	// 是否开启L7层防护，固定值为fasle
	EnableL7 bool `json:"enable_L7"`

	// HTTP请求数分段ID，固定值为1
	HttpRequestPosId int64 `json:"http_request_pos_id"`

	// 流量分段ID，取值范围：1：10M;2：30M;3：50M;4：70M;5：100M;6：150M;7：200M;8：250M;9：300M;10：500M;11：800M;88：1000M;99：默认防护。
	TrafficPosId int64 `json:"traffic_pos_id"`
}

func (o OpenAntiDDosServiceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenAntiDDosServiceRequestBody struct{}"
	}

	return strings.Join([]string{"OpenAntiDDosServiceRequestBody", string(data)}, " ")
}

type OpenAntiDDosServiceRequestBodyAppTypeId struct {
	value int64
}

type OpenAntiDDosServiceRequestBodyAppTypeIdEnum struct {
	E_0 OpenAntiDDosServiceRequestBodyAppTypeId
	E_1 OpenAntiDDosServiceRequestBodyAppTypeId
}

func GetOpenAntiDDosServiceRequestBodyAppTypeIdEnum() OpenAntiDDosServiceRequestBodyAppTypeIdEnum {
	return OpenAntiDDosServiceRequestBodyAppTypeIdEnum{
		E_0: OpenAntiDDosServiceRequestBodyAppTypeId{
			value: 0,
		}, E_1: OpenAntiDDosServiceRequestBodyAppTypeId{
			value: 1,
		},
	}
}

func (c OpenAntiDDosServiceRequestBodyAppTypeId) Value() int64 {
	return c.value
}

func (c OpenAntiDDosServiceRequestBodyAppTypeId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenAntiDDosServiceRequestBodyAppTypeId) UnmarshalJSON(b []byte) error {
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
