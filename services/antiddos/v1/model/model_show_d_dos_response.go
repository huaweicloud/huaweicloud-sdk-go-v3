package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDDosResponse Response Object
type ShowDDosResponse struct {

	// 是否开启L7层防护，固定值为fasle
	EnableL7 *bool `json:"enable_L7,omitempty"`

	// 流量分段ID，取值范围：1：10M;2：30M;3：50M;4：70M;5：100M;6：150M;7：200M;8：250M;9：300M;10：500M;11：800M;88：1000M;99：默认防护。
	TrafficPosId *int64 `json:"traffic_pos_id,omitempty"`

	// HTTP请求数分段ID，固定值为1
	HttpRequestPosId *int64 `json:"http_request_pos_id,omitempty"`

	// 清洗时访问限制分段ID，取值范围：1：10M;2：30M;3：50M;4：70M;5：100M;6：150M;7：200M;8：250M;9：300M;10：500M;11：800M;88：1000M;99：默认防护。
	CleaningAccessPosId *int64 `json:"cleaning_access_pos_id,omitempty"`

	// 应用类型ID，固定值为0
	AppTypeId      *ShowDDosResponseAppTypeId `json:"app_type_id,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowDDosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDDosResponse struct{}"
	}

	return strings.Join([]string{"ShowDDosResponse", string(data)}, " ")
}

type ShowDDosResponseAppTypeId struct {
	value int64
}

type ShowDDosResponseAppTypeIdEnum struct {
	E_0 ShowDDosResponseAppTypeId
	E_1 ShowDDosResponseAppTypeId
}

func GetShowDDosResponseAppTypeIdEnum() ShowDDosResponseAppTypeIdEnum {
	return ShowDDosResponseAppTypeIdEnum{
		E_0: ShowDDosResponseAppTypeId{
			value: 0,
		}, E_1: ShowDDosResponseAppTypeId{
			value: 1,
		},
	}
}

func (c ShowDDosResponseAppTypeId) Value() int64 {
	return c.value
}

func (c ShowDDosResponseAppTypeId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDDosResponseAppTypeId) UnmarshalJSON(b []byte) error {
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
