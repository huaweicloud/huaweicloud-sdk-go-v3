package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateFlowOutputRequestBody struct {

	// ip白名单，最大20条ip白名单
	CidrWhitelist *[]string `json:"cidr_whitelist,omitempty"`

	// 推流地址，支持ip和域名，最大值64
	Destination *string `json:"destination,omitempty"`

	Encryption *FlowSourceDecryption `json:"encryption,omitempty"`

	// 输出状态，ENABLED：启用，DISABLED：禁用
	OutputStatus *UpdateFlowOutputRequestBodyOutputStatus `json:"output_status,omitempty"`

	// 端口
	Port *int32 `json:"port,omitempty"`

	// srt-caller模式支持设置streamid
	StreamId *string `json:"stream_id,omitempty"`

	// 输出描述
	Description *string `json:"description,omitempty"`

	// 最小时延，单位毫秒，默认值2000，最小值10，最大值15000
	MinLatency *int32 `json:"min_latency,omitempty"`
}

func (o UpdateFlowOutputRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlowOutputRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFlowOutputRequestBody", string(data)}, " ")
}

type UpdateFlowOutputRequestBodyOutputStatus struct {
	value string
}

type UpdateFlowOutputRequestBodyOutputStatusEnum struct {
	ENABLED  UpdateFlowOutputRequestBodyOutputStatus
	DISABLED UpdateFlowOutputRequestBodyOutputStatus
}

func GetUpdateFlowOutputRequestBodyOutputStatusEnum() UpdateFlowOutputRequestBodyOutputStatusEnum {
	return UpdateFlowOutputRequestBodyOutputStatusEnum{
		ENABLED: UpdateFlowOutputRequestBodyOutputStatus{
			value: "ENABLED",
		},
		DISABLED: UpdateFlowOutputRequestBodyOutputStatus{
			value: "DISABLED",
		},
	}
}

func (c UpdateFlowOutputRequestBodyOutputStatus) Value() string {
	return c.value
}

func (c UpdateFlowOutputRequestBodyOutputStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateFlowOutputRequestBodyOutputStatus) UnmarshalJSON(b []byte) error {
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
