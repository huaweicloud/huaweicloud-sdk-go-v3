package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDDosStatusResponse Response Object
type ShowDDosStatusResponse struct {

	// 防护状态，可选范围：   - normal：表示正常   - configging：表示设置中   - notConfig：表示未设置   - packetcleaning：表示清洗   - packetdropping：表示黑洞
	Status         *ShowDDosStatusResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ShowDDosStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDDosStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowDDosStatusResponse", string(data)}, " ")
}

type ShowDDosStatusResponseStatus struct {
	value string
}

type ShowDDosStatusResponseStatusEnum struct {
	NORMAL         ShowDDosStatusResponseStatus
	CONFIGGING     ShowDDosStatusResponseStatus
	NOT_CONFIG     ShowDDosStatusResponseStatus
	PACKETCLEANING ShowDDosStatusResponseStatus
	PACKETDROPPING ShowDDosStatusResponseStatus
}

func GetShowDDosStatusResponseStatusEnum() ShowDDosStatusResponseStatusEnum {
	return ShowDDosStatusResponseStatusEnum{
		NORMAL: ShowDDosStatusResponseStatus{
			value: "normal",
		},
		CONFIGGING: ShowDDosStatusResponseStatus{
			value: "configging",
		},
		NOT_CONFIG: ShowDDosStatusResponseStatus{
			value: "notConfig",
		},
		PACKETCLEANING: ShowDDosStatusResponseStatus{
			value: "packetcleaning",
		},
		PACKETDROPPING: ShowDDosStatusResponseStatus{
			value: "packetdropping",
		},
	}
}

func (c ShowDDosStatusResponseStatus) Value() string {
	return c.value
}

func (c ShowDDosStatusResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDDosStatusResponseStatus) UnmarshalJSON(b []byte) error {
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
