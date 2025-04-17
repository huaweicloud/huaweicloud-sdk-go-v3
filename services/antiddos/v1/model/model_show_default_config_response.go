package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDefaultConfigResponse Response Object
type ShowDefaultConfigResponse struct {

	// 是否开启L7层防护，固定值为fasle
	EnableL7 *bool `json:"enable_L7,omitempty"`

	// 流量分段ID，取值范围：1：10M;2：30M;3：50M;4：70M;5：100M;6：150M;7：200M;8：250M;9：300M;10：500M;11：800M;88：1000M;99：默认防护。
	TrafficPosId *int64 `json:"traffic_pos_id,omitempty"`

	// HTTP请求数分段ID，固定值为1
	HttpRequestPosId *int64 `json:"http_request_pos_id,omitempty"`

	// 清洗时访问限制分段ID，取值范围：1：10M;2：30M;3：50M;4：70M;5：100M;6：150M;7：200M;8：250M;9：300M;10：500M;11：800M;88：1000M;99：默认防护。
	CleaningAccessPosId *int64 `json:"cleaning_access_pos_id,omitempty"`

	// 应用类型ID，固定值为0
	AppTypeId      *ShowDefaultConfigResponseAppTypeId `json:"app_type_id,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ShowDefaultConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDefaultConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowDefaultConfigResponse", string(data)}, " ")
}

type ShowDefaultConfigResponseAppTypeId struct {
	value int64
}

type ShowDefaultConfigResponseAppTypeIdEnum struct {
	E_0 ShowDefaultConfigResponseAppTypeId
	E_1 ShowDefaultConfigResponseAppTypeId
}

func GetShowDefaultConfigResponseAppTypeIdEnum() ShowDefaultConfigResponseAppTypeIdEnum {
	return ShowDefaultConfigResponseAppTypeIdEnum{
		E_0: ShowDefaultConfigResponseAppTypeId{
			value: 0,
		}, E_1: ShowDefaultConfigResponseAppTypeId{
			value: 1,
		},
	}
}

func (c ShowDefaultConfigResponseAppTypeId) Value() int64 {
	return c.value
}

func (c ShowDefaultConfigResponseAppTypeId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDefaultConfigResponseAppTypeId) UnmarshalJSON(b []byte) error {
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
