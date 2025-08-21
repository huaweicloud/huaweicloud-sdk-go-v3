package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowOutputInfoRequest Request Object
type ShowOutputInfoRequest struct {

	// true
	DataDisplay *ShowOutputInfoRequestDataDisplay `json:"data_display,omitempty"`

	// 流id
	FlowId string `json:"flow_id"`

	// 输出名称
	OutputName string `json:"output_name"`
}

func (o ShowOutputInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOutputInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowOutputInfoRequest", string(data)}, " ")
}

type ShowOutputInfoRequestDataDisplay struct {
	value string
}

type ShowOutputInfoRequestDataDisplayEnum struct {
	TRUE ShowOutputInfoRequestDataDisplay
}

func GetShowOutputInfoRequestDataDisplayEnum() ShowOutputInfoRequestDataDisplayEnum {
	return ShowOutputInfoRequestDataDisplayEnum{
		TRUE: ShowOutputInfoRequestDataDisplay{
			value: "true",
		},
	}
}

func (c ShowOutputInfoRequestDataDisplay) Value() string {
	return c.value
}

func (c ShowOutputInfoRequestDataDisplay) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowOutputInfoRequestDataDisplay) UnmarshalJSON(b []byte) error {
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
